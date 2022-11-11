import * as React from "react";
import { Bar, BarChart, ResponsiveContainer, XAxis, YAxis } from 'recharts';
import { Tooltip } from "react-leaflet";
import { CalendarDateTime, isSameDay, getLocalTimeZone } from '@internationalized/date';
import format from 'tinydate'

import type { Tweet } from "../types";

interface Data {
  name: string,
  value: number,
}

const findBounds = (tweets: Tweet[]): [CalendarDateTime, CalendarDateTime] => {
  let oldest, newest;
  oldest = newest = tweets[0].date;
  for (const element of tweets) {
    let curr = element.date;
    if (curr.compare(newest) > 0)
      newest = curr;
    if (curr.compare(oldest) < 0)
      oldest = curr;
  }
  return [oldest, newest];
}

enum TimeDifference {
  Day,
  Hour,
  Minutes
}

const dayFormatter = format("{MM}/{DD}/{YYYY}"),
  hourFormatter = format("{HH}"),
  minutesFormatter = format("{HH}:{mm}")

interface TweetBarsProps {
  tweets: Tweet[];
}

const TweetBars: React.FC<TweetBarsProps> = ({ tweets }) => {
  if (tweets.length == 0)
    return null

  const data = React.useMemo(() => {
    let [oldest, newest] = findBounds(tweets);
    const map: Map<string, [number, CalendarDateTime]> = new Map();
    let diff = TimeDifference.Day
    if (isSameDay(oldest, newest) && oldest.hour - newest.hour <= 2) {
      diff = TimeDifference.Hour
      if (oldest.hour == newest.hour) {
        diff = TimeDifference.Minutes
      }
    }
    for (const element of tweets) {
      let key;
      switch (diff) {
        case TimeDifference.Day:
          key = dayFormatter(element.date.toDate(getLocalTimeZone()))
          break;
        case TimeDifference.Hour:
          key = hourFormatter(element.date.toDate(getLocalTimeZone()))
          break;
        case TimeDifference.Minutes:
          key = minutesFormatter(element.date.toDate(getLocalTimeZone()))
          break;
      }
      let val = (map.get(key) || [0])[0] + 1;
      map.set(key, [val, element.date]);
    }

    return Array.from(map.keys()).reduce<(Data & { date: CalendarDateTime })[]>((prev, name) => [...prev, {
      name,
      value: map.get(name)![0], date: map.get(name)![1]
    }], []).sort((a, b) => a.date.compare(b.date));
  }, [tweets])

  return (
    <div className="w-full h-full">
      <ResponsiveContainer width="100%" height="100%">
        <BarChart height={100} data={data}>
          <XAxis dataKey="name" />
          <YAxis />
          <Tooltip />
          <Bar isAnimationActive={true} dataKey="value" fill="#0284c7" />
        </BarChart>

      </ResponsiveContainer>
    </div>
  );
};

export default TweetBars;
