import * as React from "react";

import type { Tweet } from "../types";
import WordCloud from "react-d3-cloud"
import { scaleOrdinal } from 'd3-scale';
import { schemeCategory10 } from 'd3-scale-chromatic';

interface TermCouldProps {
  tweets: Tweet[];
}
interface Word {
  [key: string]: any;
  text: string;
  value: number;
}

const TermCloud: React.FC<TermCouldProps> = ({ tweets }) => {
  let obj: { [key: string]: number } = {};
  for (const element of tweets) {
    const words = element.text.split(" ");
    for (const word of words) {
      obj[word] = (obj[word] || 0) + 1;
    }
  }

  const words: Word[] = Object.keys(obj)
    .map((text) => ({ text, value: obj[text] }))
    .filter((word) => word.value > 1)
    .sort((a, b) => (a.value > b.value ? 1 : a.value < b.value ? -1 : 0))
    .slice(0, 80);

  const schemeCategory10ScaleOrdinal = scaleOrdinal(schemeCategory10);

  return (
    <div className="bg-white dark:bg-gray-900 px-5 text-sm font-light">
      <WordCloud
        data={words}
        font="Times"
        fontStyle="italic"
        fontWeight="bold"
        fontSize={(word) => Math.log2(word.value) * 50}
        spiral="rectangular"
        rotate={(word) => word.value % 360}
        padding={5}
        random={Math.random}
        fill={(d, i) => schemeCategory10ScaleOrdinal(i)}
        onWordClick={(event, d) => {
          console.log(`onWordClick: ${d.text}`);
        }}
        onWordMouseOver={(event, d) => {
          console.log(`onWordMouseOver: ${d.text}`);
        }}
        onWordMouseOut={(event, d) => {
          console.log(`onWordMouseOut: ${d.text}`);
        }}
      />
    </div>
  );
};

export default TermCloud;
