import * as React from 'react';
import { PieChart } from 'react-minimal-pie-chart';

import { Sentiment } from '../types';

function TweetCake() {
  //to edit with dataset in input
  const dataset = [
    { title: Sentiment.Anger, value: 4, color: '#E38627' },
    { title: Sentiment.Sadness, value: 3, color: '#C13C37' },
    { title: Sentiment.Fear, value: 4, color: '#6A2135' },
    { title: Sentiment.Joy, value: 6, color: '#32CD32' }
  ]
  const [selected, setSelected] = React.useState<number | undefined>(undefined);
  const [hovered, setHovered] = React.useState<number | undefined>(undefined);
  const data = dataset.map((entry, i) => {
    if (hovered === i) {
      return {
        ...entry,
        color: 'grey',
      };
    }
    return entry;
  });

  const lineWidth = 60;
  return (
    <PieChart
      data={data}
      radius={24}
      lineWidth={lineWidth}
      segmentsStyle={{ transition: 'stroke .3s', cursor: 'pointer' }}
      segmentsShift={(index) => (index === selected ? 2 : 0.5)}
      center={[50, 30]}
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

export default TweetCake;
