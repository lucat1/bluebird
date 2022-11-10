import * as React from 'react';
import { PieChart } from 'react-minimal-pie-chart';

import { Tweet, Sentiment } from '../types';

interface TweetCakeProps {
  tweets: Tweet[]
}

const TweetCake: React.FC<TweetCakeProps> = ({ tweets }) => {
  //to edit with dataset in input
  const dataset = [
    { title: Sentiment.Anger, value: 4, color: '#0c4a6e' },
    { title: Sentiment.Sadness, value: 4, color: '#a16207' },
    { title: Sentiment.Fear, value: 4, color: '#6A2135' },
    { title: Sentiment.Joy, value: 4, color: '#047857' }
  ]
  const [selected, setSelected] = React.useState<number | undefined>(undefined);
  const [hovered, setHovered] = React.useState<number | undefined>(undefined);
  const data = dataset.map((entry, i) => {
    if (hovered === i) {
      return {
        ...entry,
        color: '#0ea5e9',
      };
    }
    return entry;
  });

  const lineWidth = 60;
  return (
    <PieChart
      data={data}
      radius={48}
      lineWidth={lineWidth}
      segmentsStyle={{ transition: 'stroke .3s', cursor: 'pointer' }}
      segmentsShift={(index) => (index === selected ? 2 : 0.5)}
      center={[50, 50]}
      animate
      label={({ dataEntry }) => dataEntry.title + ' ' + Math.round(dataEntry.percentage) + '%'}
      labelPosition={100 - lineWidth / 2}
      labelStyle={{
        fontSize: 3,
        fill: '#fff',
        opacity: 0.75,
        pointerEvents: 'none',
      }}
      onClick={(_, index) => setSelected(index === selected ? undefined : index)}
      onMouseOver={(_, index) => setHovered(index)}
      onMouseOut={() => setHovered(undefined)}
    />
  );
}

export default TweetCake
