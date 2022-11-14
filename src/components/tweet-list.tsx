import * as React from "react";
import { getLocalTimeZone } from '@internationalized/date';
import format from 'tinydate'

import useStore from '../store'
import Legend from "./legend";

const dateFormatter = format("{DD}/{MM}/{YYYY} {HH}:{mm}")

const TweetList: React.FC = () => {
  const tweets = useStore(s => s.tweets)

  return (
    <>
      <div className="flex justify-center mb-4">
        <span className="dark:text-white">Found <span className="text-sky-800 dark:text-sky-600">{tweets.length || 0}</span> tweet{tweets.length > 1 && 's'}</span>
      </div>
      {tweets.map(tweet => (
        <div key={tweet.id} className="dark:bg-gray-800 p-4 my-4 rounded-lg shadow-2xl shadow-zinc-400 dark:shadow-sky-900 border dark:border-gray-600">
          <div className="flex items-center justify-between mb-4">
            <a
              className="flex space-x-4"
              href={`https://twitter.com/${tweet.user.username}`}
              target="_blank"
            >
              <img
                className="w-10 h-10 rounded-full"
                src={tweet.user.profile_image}
                alt={`${tweet.user.name}'s profile picture`}
              />
              <div className="font-medium dark:text-white">
                <div>{tweet.user.name}</div>
                <div className="text-sm text-gray-500 dark:text-gray-400">
                  @{tweet.user.username}
                </div>
              </div>
            </a>
            <span>{dateFormatter(tweet.date.toDate(getLocalTimeZone()))}</span>
          </div>
          {tweet.text}
          {tweet.sentiments && (
            <Legend sentiments={tweet.sentiments || []} />
          )}
        </div>
      ))}
    </>
  );
}

export default TweetList;
