import * as React from "react";
import { getLocalTimeZone } from "@internationalized/date";
import format from "tinydate";

import useStore from "../stores/eredita";
import Legend from "./legend";

const dateFormatter = format("{DD}/{MM}/{YYYY} {HH}:{mm}");
const dayFormatter = format("{DD}/{MM}/{YYYY}");

//const correctWord = /cane/ig
const correctDate = "14/11/2022";
//cout number of wrong and right tweets
let right = 0,
  wrong = 0;

const TweetList: React.FC = () => {
  const tweets = useStore((s) => s.tweets);
  const { ghigliottina } = useStore(s => ({ ghigliottina: s.ghigliottina, loadingGhigliottina: s.loadingGhigliottina }))
  let correctWord = /.*/g
  if(ghigliottina)
    correctWord = new RegExp(ghigliottina.word)
  return (
    <>
      <div className="flex justify-center mb-4">
        <span className="dark:text-white">
          Found{" "}
          <span className="text-sky-800 dark:text-sky-600">
            {tweets.length || 0}
          </span>{" "}
          tweet{tweets.length > 1 && "s"}
        </span>
      </div>
      {tweets.map((tweet) => {
        const isRightWord = correctWord.test(tweet.text.toUpperCase());
        isRightWord &&
        correctDate == dayFormatter(tweet.date.toDate(getLocalTimeZone()))
          ? ++right
          : ++wrong;
        return (
          <div
            key={tweet.id}
            className={`${
              isRightWord &&
              correctDate == dayFormatter(tweet.date.toDate(getLocalTimeZone()))
                ? "border border-green-600"
                : "border border-red-600"
            } dark:bg-gray-800 p-4 my-4 rounded-lg shadow-2xl shadow-zinc-400 dark:shadow-sky-900`}
          >
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
              <span>
                {dateFormatter(tweet.date.toDate(getLocalTimeZone()))}
              </span>
            </div>
            {tweet.text}
            {tweet.sentiments && <Legend sentiments={tweet.sentiments || []} />}
          </div>
        );
      })}
    </>
  );
};

export default TweetList;