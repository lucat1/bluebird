import * as React from 'react'
import { useQuery } from '@tanstack/react-query'
import fetch from '../fetch'
import type { Tweet } from '../types'

const Search: React.FC = () => {
  const { data: tweets } = useQuery([], () => fetch<Tweet[]>('search'), { suspense: true })
  const query = ''
  return (
    <>
      <form method="GET" url="/search/keyword" className="m-4 grid grid-cols-6 grap-4">
        <label htmlFor="search" className="mb-2 text-sm font-medium text-gray-900 sr-only dark:text-gray-300">Your Email</label>
        <div className="relative col-start-3 col-span-2">
          <div className="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
            <svg aria-hidden="true" className="w-5 h-5 text-gray-500 dark:text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path></svg>
          </div>
          <input type="search" id="query" name="query" className="block p-4 pl-10 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-sky-500 focus:border-sky-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-sky-500 dark:focus:border-sky-500" placeholder="Search" required={true} value={query} />
          <button type="submit" className="text-white absolute right-2.5 bottom-2.5 bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800">Search</button>
        </div>
      </form>

      {tweets?.map((tweet) => (
        <div className="grid grid-cols-6 grap-4 text-left">
          <div className="dark:bg-gray-800 p-6 rounded-lg border col-start-2 col-span-4 shadow-2xl m-4 shadow-sky-900 shadow-grey-300 focus:ring-sky-500 focus:border-sky-500 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-sky-500 dark:focus:border-sky-500">
            <div className="flex items-center justify-between mb-4">
              <a className="flex space-x-4" href="https://twitter.com/{{this.User.Username}}" target="_blank">
                <img className="w-10 h-10 rounded-full" src={tweet.user.profile_image} alt={`${tweet.user.name}'s profile picture`} />
                <div className="font-medium dark:text-white">
                  <div>{tweet.user.name}</div>
                  <div className="text-sm text-gray-500 dark:text-gray-400">@{tweet.user.username}</div>
                </div>
              </a>
              <a className="flex space-x-4" href="https://twitter.com/{{this.User.Username}}/status/{{this.ID}}" target="_blank">
                <button className="text-white bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800">Go to tweet</button>
              </a>
            </div>
            {tweet.text}
          </div>
        </div>
      ))
      }
    </>
  )
}

export default Search
