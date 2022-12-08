
import create from "zustand";
import type { DateRange } from "@react-types/datepicker";
import { now } from "@internationalized/date";

import fetch from "../fetch";
import { convert } from "./store";
import { Search, Tweet } from "../types";

export enum QueryType {
  Keyword = "keyword",
  User = "user",
}

export interface Query {
  type: QueryType;
  query: string;
  timeRange: DateRange;
}

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
    query: "#fantacitorio",
    timeRange: {
      start: now("utc")
        .subtract({
          days: 1,
        })
        .set({ hour: 18, minute: 0 }),
      end: now("utc").subtract({ days: 1 }).set({ hour: 21, minute: 0 }),
    },
  },
  loading: true,
  tweets: [],

});

const searchURL = (url: string, { type, query, timeRange }: Query): string => {
  if (!type || !query) return url;

  let base = `${url}?type=${type}&query=${encodeURIComponent(
    query
  )}&amount=100`;
  if (timeRange) {
    const start = timeRange.start.toDate("utc").toISOString();
    const end = timeRange.end.toDate("utc").toISOString();
    base += `&startTime=${start}&endTime=${end}`;
  }
  return base;
};

const store = create<State & Actions>((set, get) => ({
  ...getInitialState(),

  reset: () => set(getInitialState()),
  clearTweets: () => set({ ...get(), tweets: [] }),
  fetch: async (query: Query) => {
    set({ ...getInitialState(), query });
    const req = await fetch<Search>(searchURL("search", query));
    const tweets = req.tweets.map(convert);
    set({ ...get(), loading: false, tweets });

  },
}));

export default store;