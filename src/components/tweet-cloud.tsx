import * as React from "react";
import shallow from "zustand/shallow";
import WordCloud from "react-d3-cloud";
import { scaleOrdinal } from "d3-scale";
// @ts-ignore
import { schemeCategory10 } from "d3-scale-chromatic";

import useStore from "../stores/store";

interface Word {
  [key: string]: any;
  text: string;
  value: number;
}

const blacklist = [
  "il",
  ".",
  "o",
  "e",
  "...",
  "se",
  "tu",
  "io",
  "noi",
  "lui",
  "i",
  "del",
  "le",
  "alle",
  ";",
  ":",
  "/",
  "da",
  "alla",
  "?",
  "!",
  "la",
  "gli",
  "lo",
  "l'",
  "un",
  "una",
  "uno",
  "dei",
  "della",
  "delle",
  "dello",
  "di",
  "e",
  "ed",
  "a",
  "ad",
  "tra",
  "in",
  "con",
  "su",
  "per",
  "fra",
  "è",
  "non",
  "ha",
  "si",
  "no",
  "al",
  "ma",
  "che",
];

const TweetCloud: React.FC = () => {
  const texts = useStore((s) => s.tweets.map((t) => t.text), shallow);
  const words = React.useMemo(() => {
    let obj: { [key: string]: number } = {};
    for (const text of texts) {
      const words = text.split(/,| /);
      for (const word of words) {
        obj[word] = (obj[word] || 0) + 1;
      }
    }
    const words = Object.keys(obj)
      .map((text) => ({ text, value: obj[text] }))
      .filter(
        (word) =>
          !!word.text.trim() &&
          word.value > 1 &&
          !blacklist.includes(word.text.toLowerCase()) &&
          !word.text.includes("#") &&
          !(word.text.length == 1) &&
          !/^\d/.test(word.text)
      )
      .sort((a, b) => (a.value < b.value ? 1 : a.value > b.value ? -1 : 0))
      .slice(0, 80);
    // const tot = words.reduce((acc, entry) => acc + entry.value, 0);
    const max = words[0].value;
    return words.map((word) => ({
      ...word,
      value: word.value / max,
    })) as Word[];
  }, [texts]);

  const schemeCategory10ScaleOrdinal = scaleOrdinal(schemeCategory10);

  return (
    <div className="bg-white dark:bg-gray-900 px-5 text-sm font-light overflow-hidden [&>*]:h-full [&>*]:flex [&>*]:justify-center [&>*>*]:h-full">
      <WordCloud
        data={words}
        width={600}
        height={340}
        fontWeight="bold"
        fontSize={(word) => Math.max(20, word.value * 80)}
        spiral="rectangular"
        rotate={() => Math.random() * 180 - 90}
        padding={5}
        random={Math.random}
        fill={(_: any, i: any) => schemeCategory10ScaleOrdinal(i)}
      />
    </div>
  );
};

export default TweetCloud;
