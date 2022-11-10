import * as React from "react";
import { useQuery } from "@tanstack/react-query";
import type { DateRange } from "@react-types/datepicker";
import { parseDateTime, getLocalTimeZone } from '@internationalized/date';
import format from 'tinydate'

import { Search, RawTweet, Tweet } from '../types'
import fetch from "../fetch";

export interface TweetProps {
  type: string;
  query: string;
  timeRange?: DateRange
}

const dateFormatter = format("{YYYY}/{MM}/{DD} {HH}:{mm}")

const url = ({ type, query, timeRange }: TweetProps): string => {
  if (!type || !query) return `search`

  let base = `search?type=${type}&query=${encodeURIComponent(query)}&amount=50`
  if (timeRange) {
    const start = timeRange.start.toDate(getLocalTimeZone()).toISOString()
    const end = timeRange.end.toDate(getLocalTimeZone()).toISOString()
    base += `&startTime=${start}&endTime=${end}`
  }
  return base
}

const convert = (raw: RawTweet): Tweet => ({ ...raw, date: parseDateTime(raw.created_at.slice(0, -1)) })

const TweetList: React.FC<TweetProps> = (props) => {
  const { data } = useQuery(
    ["search", props],
    () =>
      fetch<Search>(url(props)),
    { suspense: true }
  );
  const tweets = data!.tweets.map(convert)

  return (
    <>
      <div className="flex justify-center mb-4">
        <span className="dark:text-white">Found <span className="text-sky-800 dark:text-sky-600">{data?.tweets.length || 0}</span> tweets</span>
      </div>
      {tweets.map((tweet) => (
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
        </div>
      ))}
    </>
  );
};

export default TweetList;
