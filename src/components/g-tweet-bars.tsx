import * as React from "react";
import { Bar, BarChart, ResponsiveContainer, XAxis, YAxis } from "recharts";
import { Tooltip } from "react-leaflet";
import {
  CalendarDateTime,
  // isSameDay,
  getLocalTimeZone,
} from "@internationalized/date";
import format from "tinydate";

import useStore from "../stores/eredita";
// import type { Tweet } from "../types";

interface Data {
  name: string;
  right: number;
  wrong: number;
}

// const findBounds = (tweets: Tweet[]): [CalendarDateTime, CalendarDateTime] => {
//   let oldest, newest;
//   oldest = newest = tweets[0].date;
//   for (const element of tweets) {
//     let curr = element.date;
//     if (curr.compare(newest) > 0) newest = curr;
//     if (curr.compare(oldest) < 0) oldest = curr;
//   }
//   return [oldest, newest];
// };

enum TimeDifference {
  Day,
  Hour,
  Minutes,
}

// const dayFormatter = format("{DD}/{MM}/{YYYY}"),
// hourFormatter = format("{HH}"),
const minutesFormatter = format("{HH}:{mm}");

const GhigliottinaBars: React.FC = () => {
  const { tweets } = useStore((s) => ({
    tweets: s.tweets,
  }));

  const data = React.useMemo(() => {
    if (tweets.length == 0) return null;
    // let [oldest, newest] = findBounds(tweets);
    const map: Map<string, [number, number, CalendarDateTime]> = new Map();
    let diff = TimeDifference.Minutes;
    // if (isSameDay(oldest, newest) && oldest.hour - newest.hour <= 2) {
    //   diff = TimeDifference.Hour;
    //   if (oldest.hour == newest.hour) {
    //     diff = TimeDifference.Minutes;
    //   }
    // }

    for (const element of tweets) {
      let key;
      switch (diff) {
        // case TimeDifference.Day:
        //   key = dayFormatter(element.date.toDate(getLocalTimeZone()));
        //   break;
        // case TimeDifference.Hour:
        //   key = hourFormatter(element.date.toDate(getLocalTimeZone()));
        //   break;
        case TimeDifference.Minutes:
          key = minutesFormatter(element.date.toDate(getLocalTimeZone()));
          break;
      }
      let mapElement = map.get(key) || [0, 0, element.date];
      console.log(mapElement);
      element.rightWord ? mapElement[0]++ : mapElement[1]++;
      map.set(key, mapElement);
    }

    return Array.from(map.keys())
      .reduce<(Data & { date: CalendarDateTime })[]>(
        (prev, name) => [
          ...prev,
          {
            name,
            right: map.get(name)![0],
            wrong: map.get(name)![1],
            date: map.get(name)![2],
          },
        ],
        []
      )
      .sort((a, b) => a.date.compare(b.date));
  }, [tweets]);

  if (data == null) return null;

  return (
    <div className="w-full h-full">
      <ResponsiveContainer width="100%" height="100%">
        <BarChart height={100} data={data}>
          <XAxis dataKey="name" />
          <YAxis />
          <Tooltip />
          <Bar isAnimationActive={true} dataKey="right" fill="#0284c7" />
          <Bar isAnimationActive={true} dataKey="wrong" fill="#d9381e" />
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default GhigliottinaBars;
