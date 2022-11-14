import * as React from "react";
import shallow from "zustand/shallow";
import WordCloud from "react-d3-cloud"
import { scaleOrdinal } from 'd3-scale';
// @ts-ignore
import { schemeCategory10 } from 'd3-scale-chromatic';

import useStore from '../store'

interface Word {
  [key: string]: any;
  text: string;
  value: number;
}

const TweetCloud: React.FC = () => {
  const blacklist = ["il", "la", "gli", "lo", "l'", "un", "una", "uno", "dei", "delle", "dello", "di", "e", "ed", "a", "ad", "tra", "in", "con", "su", "per", "fra"]
  const texts = useStore(s => s.tweets.map(t => t.text), shallow)
  const words = React.useMemo(() => {
    let check = false
    let obj: { [key: string]: number } = {};
    for (const text of texts) {
      const words = text.split(" ");
      for (const word of words) {
        for(const forbid of blacklist){
          if(word == forbid.toUpperCase() || word == forbid || word == (forbid.charAt(0).toUpperCase() + forbid.slice(1).toLowerCase()) || word.charAt(0) == "@"){
            check = true
            break
          }
        }
        if(!check){
          obj[word] = (obj[word] || 0) + 1;
        }
        else{
          check = false
        }
      }
    }

    return Object.keys(obj)
      .map((text) => ({ text, value: obj[text] }))
      .filter((word) => word.value > 1)
      .sort((a, b) => (a.value > b.value ? 1 : a.value < b.value ? -1 : 0))
      .slice(0, 80) as Word[];
  }, [texts])

  const schemeCategory10ScaleOrdinal = scaleOrdinal(schemeCategory10);

  return (
    <div className="bg-white dark:bg-gray-900 px-5 text-sm font-light overflow-hidden [&>*]:h-full [&>*]:flex [&>*]:justify-center [&>*>*]:h-full">
      <WordCloud
        data={words}
        width={600}
        height={340}
        fontWeight="bold"
        fontSize={(word) => Math.log2(word.value) * 50}
        spiral="rectangular"
        rotate={(word) => word.value % 360}
        padding={5}
        random={Math.random}
        fill={(_: any, i: any) => schemeCategory10ScaleOrdinal(i)}
      />
    </div>
  );
};

export default TweetCloud;
