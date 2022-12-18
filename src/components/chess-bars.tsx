import * as React from "react";
import { Bar, BarChart, ResponsiveContainer, XAxis, YAxis } from "recharts";
import { Tooltip } from "react-leaflet";

import useChess from "../stores/chess";

interface Data {
  name: string;
  value: number;
}

const TweetBars: React.FC = () => {
  const moves = useChess((s) => s.moves);

  const data = React.useMemo(() => {
    if (moves == null || moves.length == 0) return null;

    let data: Data[];
    data = [];

    for (let key in moves) {
      data.push({ name: key, value: moves[key] });
    }

    data = data.sort((a, b) => a.value - b.value);
    return data;
  }, [moves]);

  if (data == null || data.length == 0) return null;

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
