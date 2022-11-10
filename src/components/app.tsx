import * as React from "react";
import { OverlayContainer } from "@react-aria/overlays";
import { ErrorBoundary } from 'react-error-boundary'
import { now, getLocalTimeZone } from '@internationalized/date';

import Loading from "./loading";
import Navbar from "./navbar";
import Error from './error';
import TermCloud from "./term-cloud";
import Search, { searchTypes } from "../search";
import { queryClient } from "../main";
import TweetList from "./tweet-list";
import type { TweetProps } from "./components/tweet-list";
import TweetCake from "./tweet-cake";

const App: React.FC<React.PropsWithChildren<{}>> = ({ children }) => {
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
        <div className="overflow-hidden lg:flex-1 grid grid-rows-none lg:grid-cols-[2fr_1fr] lg:grid-rows-[min-content_auto] lg:gap-x-4 mx-2 lg:mx-4">
          <React.Suspense fallback={
            <div className="lg:row-span-2 flex items-center justify-center">
              <Loading />
            </div>
          }>
            <div className="row-start-2 lg:row-span-2 lg:row-start-1 grid lg:grid-row-[min-content_auto]">
              <div className="grid lg:grid-cols-[1fr_2fr] lg:row-start-1">
                <div className="lg:col-start-1 flex items-center justify-center">
                  <TweetCake />
                </div>
                <div className="lg:col-start-2"> barre</div>
              </div>
              <div className="lg:row-start-2"> TermCloud</div>
              <div className="lg:row-start-3"> bottone mappa</div>
            </div>
            <div className="lg:row-start-2 lg:col-start-2 overflow-auto">
              {props.query != "" && <TweetList {...props} />}
            </div>
          </React.Suspense>
          <div className="row-start-1 lg:col-start-2 flex justify-center">
            <Search values={props} onSubmit={setProps} />
          </div>
        </div>

      </main>
    </OverlayContainer>
  );
}
export default App;
