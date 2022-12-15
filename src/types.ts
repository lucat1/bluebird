import { CalendarDateTime } from "@internationalized/date";

export interface SentimentSearch {
  sentiments: Sentiments;
}

export enum SentimentLabel {
  Anger = "anger",
  Sadness = "sadness",
  Fear = "fear",
  Joy = "joy",
}

export interface Sentiment {
  label: SentimentLabel;
  score: number;
}

export type Sentiments = [Sentiment, Sentiment, Sentiment, Sentiment];

export interface Search {
  tweets: RawTweet[];
  cached: number;
}

export interface RawTweet {
  id: string;
  text: string;
  user: User;
  created_at: string;
  geo?: Geo;
  sentiments?: Sentiments;
}

export interface Tweet extends RawTweet {
  date: CalendarDateTime;
}

export interface Geo {
  type: string;
  id: string;
  coordinates: [number, number] | [number, number, number, number];
}

export interface User {
  id: string;
  name: string;
  username: string;
  profile_image: string;
}

export interface Ghigliottina {
  word: string;
  podium: GhigliottinaPodium;
}

export interface GhigliottinaPodium {
  first: GhigliottinaWinnder;
  second: GhigliottinaWinnder;
  third: GhigliottinaWinnder;
}

export interface GhigliottinaWinnder {
  username: string;
  time: string;
}

export enum ChessMessageType {
  Match = "match",
  Start = "start",
  Tweets = "tweets",
  Move = "move",
  Forfeit = "forfeit",
}

export interface IncomingMessage<T> {
  type: ChessMessageType;
  data: T;
}

export interface OutgoingMessage {
  type: ChessMessageType;
  data: string;
}

export interface Match {
  code: string;
  duration: string;
  ends_at: string;
  game: string;
  tweets: Tweet[] | null;
  forfeited: boolean;
}

export interface Politician {
  id: number;
  name: string;
  surname: string;
  points: number;
  //- : number; NPosts
  average: number;
  best_single_score: number;
  //- : number; LastUpdated
}

export interface PoliticiansScoreboard {
	politicians: Politician[];
  best_climber: Politician;
  best_average: Politician;
  best_single_score: Politician;
}

export interface Team {
  username: string;
  picture_url: string;
}