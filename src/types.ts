export interface Tweet {
  id: string
  text: string
  user: User
  geo?: Geo
}

export interface Geo {
  coordinates: [number, number],
}

export interface User {
  id: string
  name: string
  username: string
  profile_image: string
}
