import * as React from "react";
import { useQuery } from "@tanstack/react-query";
import fetch from "../fetch";
import { getLocalTimeZone } from '@internationalized/date';

import type { DateRange } from "@react-types/datepicker";
import type { Search, Tweet } from "../types";
import ReactWordcloud, { Word } from 'react-wordcloud';

const words = [
  {
    text: 'told',
    value: 64,
  },
  {
    text: 'mistake',
    value: 11,
  },
  {
    text: 'thought',
    value: 16,
  },
  {
    text: 'bad',
    value: 17,
  },
]

interface TermCouldProps {
  tweets: Tweet[]
}

const TermCloud: React.FC<TermCouldProps> = ({tweets}) => {
  //tweets[0].text.split(' ')
  let obj: { [key: string]: number } = {}
  for (const element of tweets) {
    const words = element.text.split(' ')
    for(const word of words){
      obj[word] = (obj[word] || 0) + 1
    }
  }

  const words: Word[] = Object.keys(obj)
    .map(text => ({ text, value: obj[text] }))
    .filter(word =>word.value > 1)
    .sort((a,b) => a.value > b.value ? 1 : (a.value<b.value ? -1 : 0)).slice(0, 80)
  //const words = React.useMemo(() => [], [tweets])

  return(
    <div className="bg-white dark:bg-gray-900 px-5 py-3  text-sm font-light">
      <ReactWordcloud options={{
        fontFamily: 'courier new',
        fontSizes: [20, 60],
      }} words={words} />

      </div>
  );
}

export default TermCloud;
