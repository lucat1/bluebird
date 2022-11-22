import { CalendarDateTime } from '@internationalized/date';

export interface SentimentSearch {
  sentiments: Sentiments
}

export enum SentimentLabel {
  Anger = 'anger',
  Sadness = 'sadness',
  Fear = 'fear',
  Joy = 'joy'
}

export interface Sentiment {
  label: SentimentLabel
  score: number
}

export type Sentiments = [Sentiment, Sentiment, Sentiment, Sentiment]

export interface Search {
  tweets: RawTweet[]
  cached: number
}

export interface RawTweet {
  id: string
  text: string
  user: User
  created_at: string
  geo?: Geo
  sentiments?: Sentiments
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

export enum ChessState {
  IDLE = 'idle',
  WAITING = 'waiting',
  MOVING = 'moving',
  LOST = 'lost'
}

export interface Chess {
  state: ChessState
  fen: string
  turn: boolean
  ends_at: string
  code: string
}
