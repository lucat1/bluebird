import * as React from "react";
//import {CalendarDate,parseDate} from '@internationalized/date';
import type { Tweet } from "../types";
import {Bar, BarChart,ResponsiveContainer, XAxis, YAxis} from 'recharts';
import { Tooltip } from "react-leaflet";
import {CalendarDateTime,parseDateTime} from '@internationalized/date';

interface BarGraphProps {
  tweets: Tweet[];
}

interface Data {
  name: string,
  value: number,
}

function findMinAndMax(tweets:Tweet[]){
  let curr , oldest, newest; 
  oldest=newest=new Date(tweets[0].created_at);

  for(const element of tweets){
      curr=new Date(element.created_at);
      if(curr>newest)
        newest=curr;
      if(curr<oldest)
        oldest=curr;
  }  
  console.log("Oldest:",oldest.toString()," Newest:",newest.toString());
  console.log(parseDateTime(oldest.toISOString()));
  return [oldest,newest];

}


const BarGraph: React.FC<BarGraphProps> = ({ tweets }) => {
  let obj: { [key: string]: number } = {};
  let [oldest,newest]=findMinAndMax(tweets);
  for (const element of tweets) {
    obj[element.created_at]=(obj[element.created_at]||0)+1;
  }

  let data : Data[] = Object.keys(obj).reduce<Data[]>( (prev,name)=>[ ...prev, { name, value:obj[name] }],[]);


  return (
    <ResponsiveContainer className="bg-white dark:bg-gray-900 px-5 py-3 text-sm font-light">
      <BarChart width={500} height={200} data={data}>
      <XAxis dataKey="name" />
      <YAxis />
      <Tooltip />
      <Bar dataKey="value" fill="#8884d8" />
      </BarChart>

    </ResponsiveContainer>
  );
};

export default BarGraph;
