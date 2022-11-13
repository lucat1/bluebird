import * as React from "react";
import { OverlayContainer } from "@react-aria/overlays";
import { ErrorBoundary } from 'react-error-boundary'
import shallow from 'zustand/shallow'

import useStore from '../store'
import Error from './error';
import Loading from "./loading";
import Navbar from "./navbar";
import Search from "./search";
import TweetCloud from "./tweet-cloud";
import TweetList from "./tweet-list";
import TweetBars from "./tweet-bars";
import TweetCake from "./tweet-cake";

const App: React.FC = () => {
  const { query, loading, reset } = useStore(s => ({ query: s.query, loading: s.loading, reset: s.reset }), shallow)

  return (
    <OverlayContainer>
      <main className="w-screen h-screen flex flex-col overflow-auto lg:overflow-hidden dark:bg-gray-900 dark:text-gray-200 dark:text-white">
        <Navbar />
        <div className="lg:flex-1 lg:overflow-auto grid grid-rows-[auto_auto_auto] grid-cols-1 lg:grid-cols-[auto_min-content] lg:grid-rows-[auto_auto] lg:gap-x-4 px-2 lg:mx-4">
          <ErrorBoundary FallbackComponent={Error} onReset={reset}>
            {query.query != "" ? loading ? (
              <div className="lg:row-span-2 flex items-center justify-center">
                <Loading />
              </div>
            ) : (
              <>
                <div className="row-start-2 lg:row-start-1 lg:row-span-2 col-span-1 flex flex-col overflow-auto xl:overflow-hidden lg:flex-1">
                  <div className="lg:p-4 flex flex-col xl:flex-row flex-initial xl:h-1/2 lg:overflow-none">
                    <div className="flex items-center justify-center aspect-square p-8 lg:p-0">
                      <TweetCake />
                    </div>
                    <div className="flex items-center justify-center aspect-video p-8 lg:p-0">
                      <TweetBars />
                    </div>
                  </div>
                  <div className="lg:p-4 flex flex-col flex-initial xl:h-1/2 lg:overflow-none">
                    <TweetCloud />
                  </div>
                </div>
                <div className="row-start-4 lg:row-start-2 lg:col-start-2 lg:overflow-auto">
                  <TweetList />
                </div>
              </>
            ) : null}
          </ErrorBoundary>
          <div className={`flex items-center justify-center ${query.query == ''
            ? 'row-start-1 col-start-1 row-span-2 col-span-2'
            : 'row-start-1 col-start-1 lg:col-start-2 h-fit'
            }`}>

            <Search />
          </div>
        </div>
      </main>
    </OverlayContainer >
  );
}
export default App;
