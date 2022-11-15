import * as React from "react";
import { Controller, useForm } from "react-hook-form";
import { now, getLocalTimeZone } from '@internationalized/date';
import shallow from "zustand/shallow";

import useStore, { Query, QueryType } from '../store'
import DateRangePicker from './date-range-picker'

const Search: React.FC = () => {
  const { query, fetch } = useStore(s => ({ query: s.query, fetch: s.fetch }), shallow)
  const {
    register,
    control,
    handleSubmit,
    setValue,
    formState: { errors },
  } = useForm<Query>({ defaultValues: query });
  React.useEffect(() => {
    setValue('query', query.query)
  }, [query, setValue])

  return (
    <>
      <div className="flex flex-col items-center">
        <form
          onSubmit={handleSubmit(fetch)}
          className="flex flex-col items-center justify-center md:flex-row mb-4 max-w-4xl dark:text-white"
        >
          <div className="flex items-center mb-4 md:mb-0">
            <label
              htmlFor="type"
              className="sr-only block mb-2 text-sm font-medium text-gray-900 dark:text-gray-400"
            >
              Search type
            </label>
            <select
              id="type"
              className="w-32 bg-gray-50 hover:border-gray-400 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-sky-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              {...register("type", { required: true })}
            >
              {Object.values(QueryType).map((type) => (
                <option key={type}>{type}</option>
              ))}
            </select>
            {errors.type?.message && (
              <label className="red">{errors.type?.message}</label>
            )}
          </div>

          <div className="relative md:ml-4">
            <div className="flex absolute inset-y-0 left-0  items-center pl-3 pointer-events-none">
              <svg
                aria-hidden="true"
                className="w-5 h-5 text-gray-500 dark:text-gray-400"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  strokeWidth="2"
                  d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                ></path>
              </svg>
            </div>
            <label
              htmlFor="query"
              className="sr-only block mb-2 text-sm font-medium text-gray-900 dark:text-gray-400"
            >
              Query
            </label>
            <input
              id="query"
              type="search"
              className="block p-4 pl-10 hover:border-gray-400   text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-sky-500 focus:border-sky-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-sky-500 dark:focus:border-sky-500"
              placeholder="Search"
              {...register("query", { required: true })}
            />
            <div className="relative">
              <button
                type="submit"
                className="text-white absolute hover:bg-sky-700  right-2.5 bottom-2.5 bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800"
              >
                Search
              </button>
            </div>
          </div>
          {errors.query && <label className="text-red-500">A query is required</label>}
        </form>

        <div className="mb-4">
          <Controller
            control={control}
            name='timeRange'
            rules={{
              validate: range => {
                if (!range) return true;
                return range.end.compare(now(getLocalTimeZone())) <= 0
              }
            }}
            render={({ field: { onChange, value } }) => (
              <DateRangePicker
                label="Between dates"
                granularity="minute"
                hourCycle={24}
                hideTimeZone
                onChange={onChange}
                value={value}
              />
            )} />
          {errors.timeRange && <label className="text-red-500">Cannot pick a date in the future</label>}
        </div>
      </div>
    </>
  );
};

export default Search;
