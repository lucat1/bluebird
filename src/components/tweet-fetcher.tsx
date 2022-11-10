import * as React from 'react'
import { useQuery } from "@tanstack/react-query";
import { parseDateTime, getLocalTimeZone } from '@internationalized/date';
import type { DateRange } from "@react-types/datepicker";

import fetch from "../fetch";
import type { Search, RawTweet, Tweet } from '../types'

export interface TweetProps {
  type: string;
  query: string;
  timeRange?: DateRange
}

interface TweetFetcherProps {
  render: (tweets: Tweet[]) => JSX.Element
}

const convert = (raw: RawTweet): Tweet => ({ ...raw, date: parseDateTime(raw.created_at.slice(0, -1)) })

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

const TweetFetcher: React.FC<TweetFetcherProps & TweetProps> = ({ render, ...props }) => {
  const { data } = useQuery(
    ["search", props],
    () =>
      fetch<Search>(url(props)),
    { suspense: true }
  );
  return render(data!.tweets.map(convert))
}

export default TweetFetcher
