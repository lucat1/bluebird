export interface Search {
  tweets: Tweet[]
  cached: number
}

export enum Sentiment {
  Anger = 'anger',
  Sadness = 'sadness',
  Fear = 'fear',
  Joy = 'joy'
}

export interface Tweet {
  id: string
  text: string
  user: User
  created_at: string
  geo?: Geo
  sentiment: Sentiment
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
