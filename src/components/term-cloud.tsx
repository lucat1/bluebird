import * as React from "react";

import type { Tweet } from "../types";
import ReactWordcloud, { Word } from "react-wordcloud";

interface TermCouldProps {
  tweets: Tweet[];
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

  return (
    <div className="bg-white dark:bg-gray-900 px-5 py-3 text-sm font-light">
      <ReactWordcloud
        options={{
          fontSizes: [20, 60],
        }}
        words={words}
      />
    </div>
  );
};

export default TermCloud;
