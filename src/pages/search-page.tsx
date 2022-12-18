import * as React from "react";
import shallow from "zustand/shallow";

import useStore from "../stores/store";
import Loading from "../components/loading";
import Search from "../components/search";
import TweetCloud from "../components/tweet-cloud";
import TweetList from "../components/tweet-list";
import TweetBars from "../components/tweet-bars";
import TweetCake from "../components/tweet-cake";
import TweetMap from "../components/tweet-map";

const SearchPage: React.FC = () => {
  const { query, loading } = useStore(
    (s) => ({ query: s.query, loading: s.loading }),
    shallow
  );

  return (
    <div className="lg:flex-1 lg:overflow-auto grid grid-rows-[auto_auto_auto] grid-cols-1 lg:grid-cols-[auto_min-content] lg:grid-rows-[auto_auto] lg:gap-x-4 px-2 lg:mx-4">
      {query.query != "" ? (
        loading ? (
          <div className="lg:row-span-2 flex items-center justify-center">
            <Loading />
          </div>
        ) : (
          <>
            <div className="row-start-2 lg:row-start-1 lg:row-span-2 col-span-1 flex flex-col overflow-auto xl:overflow-hidden lg:flex-1">
              <div className="lg:overflow-auto ">
                <div className="lg:p-4 flex flex-col xl:flex-row flex-initial xl:h-1/2 lg:overflow-none">
                  <div className="flex items-center justify-center aspect-square p-8 lg:p-0">
                    <TweetCake />
                  </div>
                  <div className="flex items-center justify-center aspect-video pr-2 p-8 lg:p-0 overflow-hidden">
                    <TweetBars />
                  </div>
                </div>
                <div className="lg:p-4 flex flex-col flex-initial xl:h-1/2 lg:overflow-hidden">
                  <TweetCloud />
                </div>
                <div className="lg:p-4 flex flex-col flex-initial xl:h-1/2  p-3">
                  <TweetMap />
                </div>
              </div>
            </div>
            <div className="row-start-4 lg:row-start-2 lg:col-start-2 lg:overflow-auto">
              <TweetList />
            </div>
          </>
        )
      ) : null}
      <div
        className={`flex items-center justify-center ${
          query.query == ""
            ? "row-start-1 col-start-1 row-span-2 col-span-2"
            : "row-start-1 col-start-1 lg:col-start-2 h-fit"
        }`}
      >
        <Search />
      </div>
    </div>
  );
};
export default SearchPage;
