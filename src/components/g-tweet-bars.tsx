import * as React from "react";
import { Bar, BarChart, ResponsiveContainer, XAxis, YAxis } from "recharts";
import { Tooltip } from "react-leaflet";
import { CalendarDateTime } from "@internationalized/date";
import format from "tinydate";
import useStore, { Show } from "../stores/eredita";

interface Data {
  name: string;
  right: number;
  wrong: number;
}

const minutesFormatter = format("{HH}:{mm}");

const GhigliottinaBars: React.FC = () => {
  const { tweets, filter } = useStore((s) => ({
    tweets: s.tweets,
    filter: s.filter,
  }));

  const data = React.useMemo(() => {
    if (tweets.length == 0) return null;
    const map: Map<string, [number, number, CalendarDateTime]> = new Map();

    for (const element of tweets) {
      let key;
      key = minutesFormatter(element.date.toDate("UTC"));
      let mapElement = map.get(key) || [0, 0, element.date];
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
          <Bar
            isAnimationActive={true}
            dataKey="right"
            cursor={"pointer"}
            onClick={(_) => {
              filter(Show.Right);
            }}
            fill="#16a34a"
          />
          <Bar
            isAnimationActive={true}
            onClick={(_) => {
              filter(Show.Wrong);
            }}
            cursor={"pointer"}
            dataKey="wrong"
            fill="#DC2626"
          />
        </BarChart>
      </ResponsiveContainer>
    </div>
  );
};

export default GhigliottinaBars;
