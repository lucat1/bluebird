import * as React from "react";

import Tweet from './tweet'
import type { Sentiments, Tweet as TTweet } from "../types";

interface TweetListProps {
  tweets: TTweet[]
  sentimentsLoaded: Function
}


const TweetList: React.FC<TweetListProps> = ({ tweets, sentimentsLoaded }) => {
  const [sentiments, setSentiments] = React.useState('');
  return (
    <>
      <div className="flex justify-center mb-4">
        <span className="dark:text-white">Found <span className="text-sky-800 dark:text-sky-600">{tweets.length || 0}</span> tweets</span>
      </div>
      {tweets.map(tweet => (
        <Tweet key={tweet.id} {...tweet} sentimentsLoaded={sentimentsLoaded} />
      ))}
    </>
  )
}

export default TweetList;
