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
import TweetCloud from "./tweet-cloud";

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
      <main className="w-screen h-screen overflow-hidden lg:overflow-none flex flex-col dark:bg-gray-900 dark:text-gray-200 dark:text-white">
        <Navbar></Navbar>
        <div className="overflow-auto lg:overflow-hidden lg:flex-1 grid grid-rows-1 grid-cols-1 lg:grid-cols-[2fr_1fr] lg:grid-rows-[min-content_auto] lg:gap-x-4 mx-2 lg:mx-4">
          <ErrorBoundary FallbackComponent={Error}>
            <React.Suspense fallback={
              <div className="lg:row-span-2 flex items-center justify-center">
                <Loading />
              </div>
            }>
              {props.query != "" && <TweetFetcher {...props} render={tweets => (
                <>
                  <div className="row-start-2 lg:row-span-2 lg:row-start-1 grid grid-rows-1 grid-cols-1  lg:grid-row-[min-content_auto] auto-rows-fr">
                    <div className="lg:row-start-1 lg:m-4 flex flex-col lg:flex-row">
                      <div className="flex items-center justify-center aspect-square p-8 lg:p-0">
                        <TweetCake tweets={tweets} />
                      </div>
                      <div className="flex items-center justify-center aspect-video p-8 lg:p-0">
                        <TweetBars tweets={tweets} />
                      </div>
                    </div>
                    <div className="lg:row-start-2 flex items-center justify-center">
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
            : 'row-start-1 col-start-1 lg:col-start-2'
            }`}>
            <Search values={props} onSubmit={setProps} />
          </div>
        </div>

      </main>
    </OverlayContainer>
  );
}
export default App;
