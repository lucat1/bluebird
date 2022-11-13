import * as React from 'react';
import { PieChart, Pie, Cell, ResponsiveContainer, Legend } from 'recharts';
import shallow from 'zustand/shallow';

import useStore from '../store'
import { Tweet, SentimentLabel, Sentiments } from '../types';

export const colorsByLabel = {
  [SentimentLabel.Anger]: '#6A2135',
  [SentimentLabel.Sadness]: '#0c4a6e',
  [SentimentLabel.Fear]: '#047857',
  [SentimentLabel.Joy]: '#a16207'
}

const translation = {
  [SentimentLabel.Anger]: 'Rabbia',
  [SentimentLabel.Sadness]: 'Tristezza',
  [SentimentLabel.Fear]: 'Paura',
  [SentimentLabel.Joy]: 'Gioia'
}

const TweetCake: React.FC = () => {
  const { sentiments: sentimentss, len } = useStore(s => ({ sentiments: s.tweets.map(t => t.sentiments).filter((s): s is Sentiments => !!s), len: s.tweets.length }), shallow)
  const byLabel: Map<SentimentLabel, number> = new Map()
  for (const sentiments of sentimentss) {
    for (const sentiment of sentiments) {
      byLabel.set(sentiment.label, (byLabel.get(sentiment.label) || 0) + sentiment.score)
    }
  }
  const dataset = Object.values(SentimentLabel).map(sentiment => ({
    name: translation[sentiment],
    value: byLabel.get(sentiment) || 0,
    color: colorsByLabel[sentiment]
  }))

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
          label={({ value }) => { return Math.round((value / len * 100)) + '%' }}
        >
          {dataset.map((entry, index) => (
            <Cell key={`cell-${index}`} fill={entry.color} stroke={entry.color} />
          ))}
        </Pie>
        <Legend verticalAlign="bottom" />
      </PieChart>
    </ResponsiveContainer>
  );
}

export default TweetCake
