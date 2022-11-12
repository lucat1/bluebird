import * as React from "react";
import { OverlayContainer } from "@react-aria/overlays";
import { ErrorBoundary } from 'react-error-boundary'
import { now, getLocalTimeZone } from '@internationalized/date';

import Error from './error';
import Loading from "./loading";
import Navbar from "./navbar";
import Search, { searchTypes } from "../search";
import TweetFetcher from "./tweet-fetcher";
import TweetList from "./tweet-list";
import TweetBars from "./tweet-bars";
import TweetCake from "./tweet-cake";
import TweetCloud from "./term-cloud";

import type { TweetProps } from "./tweet-fetcher";

const App: React.FC<React.PropsWithChildren<{}>> = () => {
  const [props, setProps] = React.useState<TweetProps>({
    type: searchTypes[0],
    query: "",
    timeRange: {
      start: now(getLocalTimeZone()).subtract({
        days: 7
      }),
      end: now(getLocalTimeZone())
    }
  });

  return (
    <OverlayContainer>
      <main className="w-screen h-screen flex flex-col overflow-auto lg:overflow-hidden dark:bg-gray-900 dark:text-gray-200 dark:text-white">
        <Navbar></Navbar>
        <div className="lg:flex-1 lg:overflow-auto grid grid-rows-[auto_auto_auto] grid-cols-1 lg:grid-cols-[auto_min-content] lg:grid-rows-[auto_auto] lg:gap-x-4 px-2 lg:mx-4">
          <ErrorBoundary FallbackComponent={Error}>
            <React.Suspense fallback={
              <div className="lg:row-span-2 flex items-center justify-center">
                <Loading />
              </div>
            }>
              {props.query != "" && <TweetFetcher {...props} render={tweets => (
                <>
                  <div className="row-start-2 lg:row-start-1 lg:row-span-2 col-span-1 flex flex-col overflow-auto xl:overflow-hidden lg:flex-1">
                    <div className="lg:p-4 flex flex-col xl:flex-row flex-initial xl:h-1/2 lg:overflow-none">
                      <div className="flex items-center justify-center aspect-square p-8 lg:p-0">
                        <TweetCake tweets={tweets} />
                      </div>
                      <div className="flex items-center justify-center aspect-video p-8 lg:p-0">
                        <TweetBars tweets={tweets} />
                      </div>
                    </div>
                    <div className="lg:p-4 flex flex-col flex-initial xl:h-1/2 lg:overflow-none">
                      <TweetCloud tweets={tweets} />
                    </div>
                  </div>
                  <div className="row-start-4 lg:row-start-2 lg:col-start-2 lg:overflow-auto">
                    {props.query != "" && <TweetList tweets={tweets} />}
                  </div>
                </>
              )} />}
            </React.Suspense>
          </ErrorBoundary>
          <div className={`flex items-center justify-center ${props.query == ''
            ? 'row-start-1 col-start-1 row-span-2 col-span-2'
            : 'row-start-1 col-start-1 lg:col-start-2 h-fit'
            }`}>
            <Search values={props} onSubmit={setProps} />
          </div>
        </div>

      </main>
    </OverlayContainer>
  );
}
export default App;
