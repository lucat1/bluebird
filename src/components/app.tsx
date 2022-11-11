import * as React from "react";
import { OverlayContainer } from "@react-aria/overlays";
import { ErrorBoundary } from 'react-error-boundary'
import { now, getLocalTimeZone } from '@internationalized/date';

import Error from './error';
import Loading from "./loading";
import Navbar from "./navbar";
import Search, { searchTypes } from "../search";
import TweetFetcher from "./tweet-fetcher";
import TweetCloud from "./tweet-cloud";
import TweetList from "./tweet-list";
import TweetBars from "./tweet-bars";
import TweetCake from "./tweet-cake";

import type { TweetProps } from "./tweet-fetcher";
import { Sentiment, SentimentData, SentimentLabel, Sentiments } from "../types";

const defaultProps = {
  type: searchTypes[0],
  query: "",
  timeRange: {
    start: now(getLocalTimeZone()).subtract({
      days: 7
    }),
    end: now(getLocalTimeZone())
  }
}
const defaultDataset = [
  { name: SentimentLabel.Anger, value: 4, color: '#0c4a6e' },
  { name: SentimentLabel.Sadness, value: 4, color: '#a16207' },
  { name: SentimentLabel.Fear, value: 4, color: '#6A2135' },
  { name: SentimentLabel.Joy, value: 4, color: '#047857' }
]
const defaultDataset2 = [
  { name: SentimentLabel.Anger, value: 5, color: '#0c4a6e' },
  { name: SentimentLabel.Sadness, value: 5, color: '#a16207' },
  { name: SentimentLabel.Fear, value: 5, color: '#6A2135' },
  { name: SentimentLabel.Joy, value: 5, color: '#047857' }
]

const App: React.FC<React.PropsWithChildren<{}>> = () => {

  const [props, setProps] = React.useState<TweetProps>(defaultProps);

  return (
    <OverlayContainer>
      <main className="w-screen h-screen flex flex-col overflow-auto lg:overflow-hidden dark:bg-gray-900 dark:text-gray-200 dark:text-white">
        <Navbar></Navbar>
        <div className="lg:flex-1 lg:overflow-auto grid grid-rows-[auto_auto_auto] grid-cols-1 lg:grid-cols-[auto_min-content] lg:grid-rows-[auto_auto] lg:gap-x-4 px-2 lg:mx-4">
          <ErrorBoundary FallbackComponent={Error} onReset={_ => setProps(defaultProps)}>
            <React.Suspense fallback={
              <div className="lg:row-span-2 flex items-center justify-center">
                <Loading />
              </div>
            }>
              {props.query != "" && <TweetFetcher {...props} render={tweets => {
                const defaultDataSentiments = [
                  { name: SentimentLabel.Anger, sum: 0, num: 0, mean: 0 },
                  { name: SentimentLabel.Sadness, sum: 0, num: 0, mean: 0 },
                  { name: SentimentLabel.Fear, sum: 0, num: 0, mean: 0 },
                  { name: SentimentLabel.Joy, sum: 0, num: 0, mean: 0 }
                ]
                let dataSentiments = defaultDataSentiments;

                const [dataset, setDataset] = React.useState<SentimentData[]>(defaultDataset);
                const sentimentsLoaded = (sents: Sentiments) => {
                  for (let i = 0; i < sents.length; i++) {
                    for (let j = 0; j < dataSentiments.length; j++) {
                      if (dataSentiments[j].name == sents[i].label) {
                        dataSentiments[j].num++;
                        dataSentiments[j].sum += sents[i].score;
                        dataSentiments[j].mean = dataSentiments[j].sum / dataSentiments[j].num;
                        break;
                      }
                    }
                  }
                  console.log(dataSentiments)
                }

                React.useEffect(() => {
                  let cp = defaultDataset;
                  for (let j = 0; j < dataSentiments.length; j++) {
                    for (let k = 0; k < cp.length; k++) {
                      if (cp[k].name == dataSentiments[j].name) {
                        cp[k].value = Number(dataSentiments[j].mean.toFixed(4)) * 100;
                        break;
                      }
                    }
                  }
                  setDataset(cp)
                  console.log(cp)
                }, [dataSentiments]);

                React.useEffect(() => {
                  dataSentiments = defaultDataSentiments;
                }, [props.query])

                return (
                  <div className="row-start-2 lg:row-start-1 lg:row-span-2 col-span-1 flex flex-col overflow-auto xl:overflow-hidden lg:flex-1">
                    <div className="lg:p-4 flex flex-col xl:flex-row flex-initial xl:h-1/2 lg:overflow-none">
                      <div className="flex items-center justify-center aspect-square p-8 lg:p-0">
                        <TweetCake tweets={tweets} dataset={dataset} key={"tweetCake"} />
                      </div>
                      <div className="flex items-center justify-center aspect-video p-8 lg:p-0">
                        <TweetBars tweets={tweets} />
                      </div>
                    </div>
                    <div className="lg:p-4 flex flex-col flex-initial xl:h-1/2 lg:overflow-none">
                      <TweetCloud tweets={tweets} />
                    </div>
                    <div className="row-start-4 lg:row-start-2 lg:col-start-2 lg:overflow-auto">
                      {props.query != "" && <TweetList tweets={tweets} sentimentsLoaded={sentimentsLoaded} />}
                    </div>
                  </div>
                )
              }} />}
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
    </OverlayContainer >
  );
}
export default App;
