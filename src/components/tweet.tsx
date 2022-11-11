import * as React from 'react'
import { useQuery } from "@tanstack/react-query";
import { getLocalTimeZone } from '@internationalized/date';
import format from 'tinydate'

import fetch from '../fetch'
import type { Tweet as TTweet, SentimentSearch, Sentiments } from '../types'

const dateFormatter = format("{YYYY}/{MM}/{DD} {HH}:{mm}")


const Sentiment: React.FC<{ id: string, sentimentsLoaded: Function }> = ({ id, sentimentsLoaded }) => {
  const { data } = useQuery(
    ["sentiment", id],
    () => fetch<SentimentSearch>(`sentiment?id=${id}`),
    { suspense: true }
  );

  return <SentimentRenderer data={data!.sentiments} sentimentsLoaded={sentimentsLoaded} />
}

const SentimentRenderer: React.FC<{ data: Sentiments, sentimentsLoaded: Function }> = ({ data, sentimentsLoaded }) => {
  sentimentsLoaded(data)
  return (
    <span>{data.map((sentiment, i) => (
      <React.Fragment key={i}>
        <a>{sentiment.label}</a>
        <a>{sentiment.score}</a>
      </React.Fragment>
    ))}</span>
  )
}

const Tweet: React.FC<TTweet> = (tweet) => {
  return (
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
      <React.Suspense fallback={<h1>loading</h1>}>
        {tweet.sentiments ? <SentimentRenderer data={tweet.sentiments} sentimentsLoaded={tweet.sentimentsLoaded} /> : <Sentiment id={tweet.id} sentimentsLoaded={tweet.sentimentsLoaded} />}
      </React.Suspense>
    </div>
  )
}

export default Tweet
