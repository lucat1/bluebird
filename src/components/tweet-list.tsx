import * as React from "react";
import { useQuery } from "@tanstack/react-query";
import fetch from "../fetch";
import type { Search } from "../types";

export interface TweetProps {
  type: string;
  query: string;
}

const TweetList: React.FC<TweetProps> = ({ type, query }) => {
  const { data: tweets } = useQuery(
    ["search", type, query],
    () =>
      fetch<Search>(
        type && query ? `search?type=${type}&query=${encodeURIComponent(query)}&amount=50` : `search`
      ),
    { suspense: true }
  );

  console.log(tweets?.tweets.map(t => t.geo).filter(t => t != null))

  return (
    <>
      {tweets?.tweets.map((tweet) => (
        <div key={tweet.id} className="grid grid-cols-6 grap-4 text-left">
          <div className="dark:bg-gray-800 p-6 rounded-lg border col-start-2 col-span-4 shadow-2xl m-4 dark:shadow-sky-900 shadow-zinc-400 focus:ring-sky-500 focus:border-sky-500 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-sky-500 dark:focus:border-sky-500">
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
                <button className="text-white bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800">
                  Go to tweet
                </button>
              </a>
            </div>
            {tweet.text}
          </div>
        </div>
      ))}
    </>
  );
};

export default TweetList;
