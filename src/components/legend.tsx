import * as React from "react";

import { Sentiments, SentimentLabel} from "../types";
import { colorsByLabel } from "./tweet-cake";

const Legend: React.FC<{ sentiments: Sentiments }> = ({ sentiments }) => (
    <div className="flex flex-row justify-between w-full">
      {sentiments.map(sentiment => (
      <div className="flex flex-row items-center">
      <div className="text-red-900 h-3 w-4 mx-2" style={{ background: colorsByLabel[sentiment.label] }} />
      {Math.round(sentiment.score * 100)}%
    </div>
      ))}
    </div>
  )

export default Legend;