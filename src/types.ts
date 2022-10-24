export interface Tweet {
  id: string
  text: string
  user: User
}

export interface User {
  id: string
  name: string
  username: string
  profile_image: string
}
