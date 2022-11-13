import * as React from "react";

import Tweet from './tweet'
import type { Tweet as TTweet } from "../types";

interface TweetListProps {
  tweets: TTweet[]
}

const TweetList: React.FC<TweetListProps> = ({ tweets }) => (
  <>
    <div className="flex justify-center mb-4">
      <span className="dark:text-white">Found <span className="text-sky-800 dark:text-sky-600">{tweets.length || 0}</span> tweets</span>
    </div>
    {tweets.map(tweet => (
      <Tweet key={tweet.id} {...tweet} />
    ))}
  </>
);

export default TweetList;
