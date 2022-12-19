import create from "zustand";
import { now } from "@internationalized/date";

import fetch, { searchURL, Query, QueryType } from "../fetch";
import { convert } from "./store";
import {
  Search,
  Tweet,
  Politician,
  PoliticiansScoreboard,
  Team,
  Points,
} from "../types";

export interface State {
  query: Query;
  loading: boolean;
  tweets: Tweet[];
  scoreboard: PoliticiansScoreboard;
  weekly: Points;
  teams: Team[];
}

export interface Actions {
  reset(): void;
  clearTweets(): void;
  fetch(query: Query): Promise<void>;
}

const emptyPol: Politician = {
  id: 0,
  name: "",
  surname: "",
  points: 0,
  average: 0,
  best_single_score: 0,
};

const getInitialState = (): State => ({
  query: {
    type: QueryType.Keyword,
    query: "#fantacitorio",
    timeRange: {
      start: now("utc").subtract({
        days: 7,
      }),
      end: now("utc"),
    },
  },
  loading: true,
  tweets: [],
  scoreboard: {
    politicians: [],
    best_climber: emptyPol,
    best_average: emptyPol,
    best_single_score: emptyPol,
  },
  teams: [],
  weekly: {
    politicians: [],
  },
});

const store = create<State & Actions>((set, get) => ({
  ...getInitialState(),

  reset: () => set(getInitialState()),
  clearTweets: () => set({ ...get(), tweets: [] }),
  fetch: async (query: Query) => {
    set({ ...getInitialState(), query });
    const req = await fetch<Search>(searchURL("search", query));
    const tweets = req.tweets.map(convert);

    const scoreboard = await fetch<PoliticiansScoreboard>(
      "fantacitorio/scoreboard"
    );
    const { teams } = await fetch<{ teams: Team[] }>("fantacitorio/teams");
    const end_date = new Date();
    const start_time = new Date(end_date);
    start_time.setDate(start_time.getDate() - 7);
    const weekly = await fetch<Points>(
      `fantacitorio/points?startTime=${start_time.toISOString()}&endTime=${end_date.toISOString()}`
    );
    set({ ...get(), loading: false, tweets, scoreboard, teams, weekly });
  },
}));

export default store;
