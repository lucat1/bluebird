import * as React from "react";
import type { Tweet } from "../types";
import {Bar, BarChart,ResponsiveContainer, XAxis, YAxis} from 'recharts';
import { Tooltip } from "react-leaflet";
import {CalendarDateTime, isSameDay,getLocalTimeZone} from '@internationalized/date';
import format from 'tinydate'

interface BarGraphProps {
  tweets: Tweet[];
}

interface Data {
  name: string,
  value: number,
}

function findMinAndMax(tweets:Tweet[]){
  let curr , oldest, newest; 
  oldest=newest=tweets[0].date;
  for(const element of tweets){
      curr=element.date;
      if(curr.compare(newest)>0)
        newest=curr;
      if(curr.compare(oldest)<0)
        oldest=curr;


  }  
  return [oldest,newest];

}

enum TimeDifference {
  Day,
  Hour,
  Minutes
}

const dayFormatter = format("{MM}/{DD}/{YYYY}"),
  hourFormatter = format("{HH}"),
  minutesFormatter = format("{HH}:{mm}")

const BarGraph: React.FC<BarGraphProps> = ({ tweets }) => {
  let [oldest,newest]=findMinAndMax(tweets);
  const map: Map<string, [number, CalendarDateTime]> = new Map();
  let diff= TimeDifference.Day
  if(isSameDay(oldest, newest) && oldest.hour - newest.hour <= 2) {
    diff = TimeDifference.Hour
    if(oldest.hour == newest.hour) {
      diff = TimeDifference.Minutes
    }
  }
  for (const element of tweets) {
    let key;
    switch(diff) {
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

  let data = Array.from(map.keys()).reduce<(Data & {date: CalendarDateTime})[]>((prev,name)=>[ ...prev, {
    name,
    value: map.get(name)![0], date: map.get(name)![1]
  }],[]).sort((a,b) => a.date.compare(b.date));

  return (
    <div className="flex justify-center width-screen h-96">
    <ResponsiveContainer width="80%" height="50%">
      <BarChart height={100} data={data}>
      <XAxis dataKey="name" />
      <YAxis />
      <Tooltip />
      <Bar dataKey="value" fill="#0284c7" />
      </BarChart>

    </ResponsiveContainer>
    </div>
  );
};

export default BarGraph;
