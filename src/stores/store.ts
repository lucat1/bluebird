import create from "zustand";
import { parseDateTime, now } from "@internationalized/date";

import fetch, { searchURL, Query, QueryType } from "../fetch";
import { Search, RawTweet, Tweet, SentimentSearch } from "../types";

export interface State {
  query: Query;
  loading: boolean;
  tweets: Tweet[];
}

export interface Actions {
  reset(): void;
  clearTweets(): void;
  fetch(query: Query): Promise<void>;
}

const getInitialState = (): State => ({
  query: {
    type: QueryType.Keyword,
    query: "",
    timeRange: {
      start: now("utc").subtract({
        days: 7,
      }),
      end: now("utc"),
    },
  },
  loading: true,
  tweets: [],
});

export const convert = (raw: RawTweet): Tweet => ({
  ...raw,
  date: parseDateTime(raw.created_at.slice(0, -1)),
});

const store = create<State & Actions>((set, get) => ({
  ...getInitialState(),

  reset: () => set(getInitialState()),
  clearTweets: () => set({ ...get(), tweets: [] }),
  fetch: async (query: Query) => {
    set({ ...get(), loading: true, query, tweets: [] });
    const res = await fetch<Search>(searchURL("search", query));
    const tweets = res.tweets.map(convert);
    set({ ...get(), loading: false, tweets });
    for (const tweet of tweets) {
      fetch<SentimentSearch>(`sentiment?id=${tweet.id}`).then(
        ({ sentiments }) =>
          set({
            ...get(),
            tweets: get().tweets.map((t) =>
              t.id == tweet.id ? { ...t, sentiments } : t
            ),
          })
      );
    }
  },
}));

export default store;
