import * as React from 'react';
import { PieChart, Pie, Cell, ResponsiveContainer } from 'recharts';

import { Tweet, SentimentLabel, SentimentData } from '../types';

interface TweetCakeProps {
  tweets: Tweet[]
  dataset: SentimentData[]
}

const TweetCake: React.FC<TweetCakeProps> = ({ tweets, dataset }) => {
  //to edit with dataset in input

  console.log("dataset")
  console.log(dataset)

  return (
      <ResponsiveContainer width="100%" height="100%">
        <PieChart width={400} height={400}>
          <Pie
            dataKey="value"
            isAnimationActive={true}
            data={dataset}
            cx="50%"
            cy="50%"
            outerRadius={100}
            fill="#ffffff"
            label
          >
            {dataset.map((entry, index) => (
              <Cell key={`cell-${index}`} fill={entry.color} stroke={entry.color} />
            ))}
          </Pie>
        </PieChart>
      </ResponsiveContainer>
  );
}

export default TweetCake
