import * as React from 'react';
import { PieChart, Pie, Cell, ResponsiveContainer } from 'recharts';

import { Tweet, SentimentLabel } from '../types';

interface TweetCakeProps {
  tweets: Tweet[]
}

const TweetCake: React.FC<TweetCakeProps> = ({ tweets }) => {
  //to edit with dataset in input
  const dataset = [
    { name: SentimentLabel.Anger, value: 4, color: '#0c4a6e' },
    { name: SentimentLabel.Sadness, value: 4, color: '#a16207' },
    { name: SentimentLabel.Fear, value: 4, color: '#6A2135' },
    { name: SentimentLabel.Joy, value: 4, color: '#047857' }
  ]

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
