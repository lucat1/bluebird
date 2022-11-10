import * as React from "react";
import { useQuery } from "@tanstack/react-query";
import { getLocalTimeZone } from '@internationalized/date';
import type { DateRange } from "@react-types/datepicker";
import type { Search } from "../types";
import fetch from "../fetch";

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
      {tweets?.tweets.map((tweet) => (
        <div key={tweet.id} className="dark:bg-gray-800 p-4 my-4 rounded-lg shadow-2xl shadow-zinc-400 dark:shadow-sky-900 border dark:border-gray-600">
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
            <span>{new Date(tweet.created_at).toLocaleString()}</span>
          </div>
          {tweet.text}
        </div>
      ))}
    </>
  );
};

export default TweetList;
