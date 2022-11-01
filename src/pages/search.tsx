import * as React from "react";
import { Controller, useForm } from "react-hook-form";

import Loading from "../components/loading";
import TweetList, { TweetForm, TweetProps } from "../components/tweet-list";

import DatePicker from "react-datepicker";

const searchTypes = ["keyword", "user"];

const Search: React.FC = () => {
  const [props, setProps] = React.useState<TweetForm>({
    type: searchTypes[0],
    query: "",
    startTime: "",
    endTime: ""
  });
  const {
    register,
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<TweetForm>();
  const onSubmit = (data: TweetForm) => {
    let td = new Date();
    if (data.startTime && data.endTime) {
      let st = new Date(data.startTime);
      let et = new Date(data.endTime);
      if (st > td || et > td) {
        alert("Non posso vedere nel futuro");
        return
      }
      if (et < st) {
        alert("La data di fine periodo non puo' essere prima di quella di inizio periodo");
        return
      }

      data.startTime = new Date(st).toISOString();
      data.endTime = new Date(et).toISOString();
      if (data.startTime == data.endTime) {
        let next = (new Date(et).getDate() + 1)
        let toSet = new Date(et)
        toSet.setDate(next)
        data.endTime = toSet.toISOString();
      }
    } else {
      let next = (new Date(td).getDate() - 7)
      let toSet = new Date(td)
      toSet.setDate(next)
      data.startTime = toSet.toISOString();
      data.endTime = new Date().toISOString();
    }
    setProps(data);
  };

  return (
    <>
      <div className="flex justify-center">
        <form
          onSubmit={handleSubmit(onSubmit)}
          className="grid grid-cols-[auto_1fr] gap-4 my-4 max-w-4xl"
        >
          <div className="flex items-center">
            <label
              htmlFor="type"
              className="sr-only block mb-2 text-sm font-medium text-gray-900 dark:text-gray-400"
            >
              Search type
            </label>
            <select
              id="type"
              className="w-32 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
              {...register("type", { required: true })}
            >
              {searchTypes.map((type) => (
                <option key={type}>{type}</option>
              ))}
            </select>
            {errors.type?.message && (
              <label className="red">{errors.type?.message}</label>
            )}
          </div>

          <div className="relative">
            <div className="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none">
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
              className="block p-4 pl-10 text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-sky-500 focus:border-sky-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-sky-500 dark:focus:border-sky-500"
              placeholder="Search"
              {...register("query", { required: true })}
            />
            <div className="relative">
              <button
                type="submit"
                className="text-white absolute right-2.5 bottom-2.5 bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800"
              >
                Search
              </button>
            </div>
          </div>

          <Controller
            control={control}
            name='startTime'
            render={({ field }) => (
              <DatePicker
                placeholderText='Select date'
                onChange={(date: Date) => { field.onChange(date) }}
                selected={field.value}
              />
            )}
          />
          <Controller
            control={control}
            name='endTime'
            render={({ field }) => (
              <DatePicker
                placeholderText='Select date'
                onChange={(date: Date) => { field.onChange(date) }}
                selected={field.value}
              />
            )}
          />

          {errors.query && "blablabla"}
        </form>
      </div>
      <React.Suspense fallback={<Loading />}>
        {props.query != "" && <TweetList {...props} />}
      </React.Suspense>
    </>
  );
};

export default Search;
