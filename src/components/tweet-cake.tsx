import React, { useState } from 'react';
import { PieChart, pieChartDefaultProps, PieChartProps } from 'react-minimal-pie-chart';

function TweetCake() {
  //to edit with dataset in input
  const dataset=[
    { title: 'Anger', value: 4, color: '#E38627' },
    { title: 'Sadness', value: 3, color: '#C13C37' },
    { title: 'Fear', value: 4, color: '#6A2135' },
    { title: 'Joy', value: 6, color: '#32CD32'}
]
  const [selected, setSelected] = useState<number | undefined>(0);
  const [hovered, setHovered] = useState<number | undefined>(undefined);
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
  pieChartDefaultProps.radius = 30;
  return (
    <PieChart
    style={{
      fontFamily:
          '"Nunito Sans", -apple-system, Helvetica, Arial, sans-serif',
          fontSize: '8px',
        }}
        data={data}
        radius={pieChartDefaultProps.radius - 6}
        lineWidth={60}
        segmentsStyle={{ transition: 'stroke .3s', cursor: 'pointer' }}
        segmentsShift={(index) => (index === selected ? 2 : 0.5)}
        center = {[50, 30]}
        animate
        label={({ dataEntry }) => Math.round(dataEntry.percentage) + '%'}
        labelPosition={100 - lineWidth / 2}
      
      labelStyle={{
        fontSize: 3,
        fill: '#fff',
        opacity: 0.75,
        pointerEvents: 'none',
      }}
      onClick={(_, index) => {
        setSelected(index === selected ? undefined : index);
      }}
      onMouseOver={(_, index) => {
        setHovered(index);
      }}
      onMouseOut={() => {
        setHovered(undefined);
      }}
    />
  );
}

export default TweetCake;