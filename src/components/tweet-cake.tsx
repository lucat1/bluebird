import * as React from 'react';
import { PieChart, Pie, Cell, ResponsiveContainer } from 'recharts';

import useStore from '../store'
import { Tweet, SentimentLabel, Sentiments } from '../types';

const colorsByLabel = {
  [SentimentLabel.Anger]: '#0c4a6e',
  [SentimentLabel.Sadness]: '#a16207',
  [SentimentLabel.Fear]: '#6A2135',
  [SentimentLabel.Joy]: '#047857'
}

const TweetCake: React.FC = () => {
  const sentimentss = useStore(s => s.tweets.map(t => t.sentiments).filter((s): s is Sentiments => !!s))
  const byLabel: Map<SentimentLabel, number> = new Map()
  for (const sentiments of sentimentss) {
    for (const sentiment of sentiments) {
      byLabel.set(sentiment.label, (byLabel.get(sentiment.label) || 0) + sentiment.score)
    }
  }
  const dataset = Object.values(SentimentLabel).map(sentiment => ({ name: sentiment, value: Number((byLabel.get(sentiment) || 0).toFixed(2)), color: colorsByLabel[sentiment] }))

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
