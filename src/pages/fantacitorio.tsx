import * as React from "react";
import shallow from 'zustand/shallow'

import useStore from '../stores/fantacitorio'
import Loading from "../components/loading";
import Search from "../components/f-search";
import TweetList from "../components/f-tweet-list";
import Slideshow from "../components/f-slideshow";


const Fantacitorio: React.FC = () => {
  const { query, loading, fetch } = useStore(s => ( {
    fetch: s.fetch,
    query: s.query,
    loading: s.loading,}), shallow)

  React.useEffect(() => {
    fetch(query);
  }, []);

  return (
        <div className="lg:flex-1 lg:overflow-auto grid grid-rows-[auto_auto_auto] grid-cols-1 lg:grid-cols-[auto_min-content] lg:grid-rows-[auto_auto] lg:gap-x-4 px-2 lg:mx-4">
            {query.query != "" ? loading ? (
              <div className="lg:row-span-2 flex items-center justify-center">
                <Loading />
              </div>
            ) : (
              <>
                <div className="row-start-2 lg:row-start-1 lg:row-span-2 col-span-1 flex flex-col overflow-auto lg:flex-1">
                  <div className="lg:p-2 flex flex-col flex-initial xl:h-1/2">
                    <div className="flex items-center justify-center">
                      Best climbers ecc...
                    </div>
                    <div className="flex items-center justify-center aspect-video lg:p-2 overflow-auto">
                      Classifica in base a tempo
                    </div>
                  </div>
                  <div className="p-2 flex items-center flex-col flex-initial xl:h-1/2  aspect-video">
                    <Slideshow />
                  </div>
                </div>
                <div className="row-start-4 lg:row-start-2 lg:col-start-2 lg:overflow-auto">
                  <TweetList />
                </div>
              </>
            ) : null}
          <div className={`flex items-center justify-center ${query.query == ''
            ? 'row-start-1 col-start-1 row-span-2 col-span-2'
            : 'row-start-1 col-start-1 lg:col-start-2 h-fit'
            }`}>

            <Search />
          </div>
        </div>
  );
}
export default Fantacitorio;
