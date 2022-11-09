import * as React from "react";
import { useQuery } from "@tanstack/react-query";
import { getLocalTimeZone } from '@internationalized/date';
import type { DateRange } from "@react-types/datepicker";
import type { Search } from "../types";
import fetch from "../fetch";
import TweetMap from "./tweet-map";
import TermCloud from "./term-cloud";
import TweetCake from "./tweet-cake"

export interface TweetProps {
  type: string;
  query: string;
  timeRange?: DateRange
}

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

const TweetList: React.FC<TweetProps> = (props) => {
  const { data: tweets } = useQuery(
    ["search", props],
    () =>
      fetch<Search>(url(props)),
    { suspense: true }
  );

  return (
    <>
      <div className="flex justify-center mb-4">
        <span className="dark:text-white">Found <span className="text-sky-800 dark:text-sky-600">{tweets?.tweets.length || 0}</span> tweets</span>
      </div>
      <TweetMap tweets={tweets?.tweets} />
      <TermCloud tweets={tweets?.tweets!} />
      <TweetCake/> 
      {tweets?.tweets.map((tweet) => (
        <div key={tweet.id} className="grid sm:grid-cols-6 grid-cols-1 text-left">
          <div className="dark:bg-gray-800 p-6 rounded-lg border sm: col-start-2 col-span-4 shadow-2xl m-4 dark:shadow-sky-900 shadow-zinc-400 focus:ring-sky-500 focus:border-sky-500 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-sky-500 dark:focus:border-sky-500">
            <div className="flex items-center justify-between mb-4">
              <a
                className="flex space-x-4"
                href="https://twitter.com/{{this.User.Username}}"
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
              <a
                className="flex space-x-4"
                href="https://twitter.com/{{this.User.Username}}/status/{{this.ID}}"
                target="_blank"
              >
              </a>
            </div>
            {tweet.text}
            <span className="block mt-4">{new Date(tweet.created_at).toLocaleString()}</span>
          </div>
        </div>
      ))}
    </>
  );
};

export default TweetList;
