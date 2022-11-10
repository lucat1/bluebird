import { CalendarDateTime } from '@internationalized/date';


export interface Search {
  tweets: RawTweet[]
  cached: number
}

export enum Sentiment {
  Anger = 'anger',
  Sadness = 'sadness',
  Fear = 'fear',
  Joy = 'joy'
}

export interface RawTweet {
  id: string
  text: string
  user: User
  created_at: string
  geo?: Geo
  sentiment: Sentiment
}

export interface Tweet extends RawTweet {
  date: CalendarDateTime
}

export interface Geo {
  type: string
  id: string
  coordinates: [number, number] | [number, number, number, number],
}

export interface User {
  id: string
  name: string
  username: string
  profile_image: string
}
