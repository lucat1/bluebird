import create from 'zustand'
import type { DateRange } from "@react-types/datepicker";
import { parseDateTime, now, getLocalTimeZone } from '@internationalized/date';

import fetch from '../fetch'
import { Search, RawTweet, Tweet, SentimentSearch } from '../types'

export enum QueryType {
  Keyword = 'keyword',
  User = 'user'
}

export interface Query {
  type: QueryType
  query: string
  timeRange: DateRange
}

export interface State {
  query: Query
  loading: boolean
  tweets: Tweet[]
}

export interface Actions {
  reset(): void
  clearTweets(): void
  fetch(query: Query): Promise<void>
}

const getInitialState = (): State => ({
  query: {
    type: QueryType.Keyword,
    query: '',
    timeRange: {
      start: now(getLocalTimeZone()).subtract({
        days: 7
      }),
      end: now(getLocalTimeZone())
    }
  },
  loading: true,
  tweets: []
})

const searchURL = ({ type, query, timeRange }: Query): string => {
  if (!type || !query) return `search`

  let base = `search?type=${type}&query=${encodeURIComponent(query)}&amount=100`
  if (timeRange) {
    const start = timeRange.start.toDate(getLocalTimeZone()).toISOString()
    const end = timeRange.end.toDate(getLocalTimeZone()).toISOString()
    base += `&startTime=${start}&endTime=${end}`
  }
  return base
}

const convert = (raw: RawTweet): Tweet => ({ ...raw, date: parseDateTime(raw.created_at.slice(0, -1)) })

const store = create<State & Actions>((set, get) => ({
  ...getInitialState(),

  reset: () => set(getInitialState()),
  clearTweets: () => set({ ...get(), tweets: [] }),
  fetch: async (query: Query) => {
    set({ ...get(), loading: true, query, tweets: [] })
    const req = await fetch<Search>(searchURL(query))
    const tweets = req.tweets.map(convert)
    set({ ...get(), loading: false, tweets })
    for (const tweet of tweets) {
      fetch<SentimentSearch>(`sentiment?id=${tweet.id}`).then(({ sentiments }) =>
        set({
          ...get(),
          tweets: get().tweets.map(t => t.id == tweet.id ? ({ ...t, sentiments }) : t)
        }))
    }
  },
}))

export default store
