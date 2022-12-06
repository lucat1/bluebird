import create from "zustand";
import type { DateRange } from "@react-types/datepicker";
import { now } from "@internationalized/date";

import fetch from "../fetch";
import { convert } from "./store";
import { Search, Tweet, SentimentSearch, Ghigliottina } from "../types";

export enum QueryType {
  Keyword = "keyword",
  User = "user",
}

export enum Show {
  All,
  Right,
  Wrong,
}

export interface gTweet extends Tweet {
  rightWord: boolean;
}

export interface Query {
  type: QueryType;
  query: string;
  timeRange: DateRange;
}

export interface State {
  query: Query;
  loading: boolean;
  tweets: gTweet[];
  loadingGhigliottina: boolean;
  ghigliottina: Ghigliottina | null;
  show: Show;
}

export interface Actions {
  reset(): void;
  clearTweets(): void;
  fetch(query: Query): Promise<void>;
  filter(choice: Show): void;
}

const getInitialState = (): State => ({
  show: Show.All,
  query: {
    type: QueryType.Keyword,
    query: "#ghigliottina",
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

  loadingGhigliottina: true,
  ghigliottina: null,
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
  filter: (choice: Show) => {
    if (choice == get().show) set({ ...get(), show: Show.All });
    else set({ ...get(), show: choice });
  },
  fetch: async (query: Query) => {
    set({ ...getInitialState(), query });
    const req = await fetch<Search>(searchURL("search", query));
    const tweets = req.tweets.map(convert);
    const gTweets = tweets.map((t) => ({ ...t, rightWord: false }));
    set({ ...get(), loading: false, tweets: gTweets });

    const diff =
      query.timeRange.end.toDate("utc").getTime() -
      query.timeRange.start.toDate("utc").getTime();
    const oneDayAndOneHour = 25 * 60 * 60 * 1000;
    if (diff <= oneDayAndOneHour) {
      set({ ...get(), loadingGhigliottina: true });
      const ghigliottina = await fetch<Ghigliottina>(
        searchURL("ghigliottina", query)
      );
      console.log(ghigliottina.word);
      const trueTweets = gTweets.map((t) => {
        if (t.text.toUpperCase().includes(ghigliottina.word))
          return { ...t, rightWord: true };
        else return { ...t, rightWord: false };
      });
      set({
        ...get(),
        loadingGhigliottina: false,
        ghigliottina,
        tweets: trueTweets,
      });
    }

    for (const tweet of get().tweets) {
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
