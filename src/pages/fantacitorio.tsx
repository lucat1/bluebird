import * as React from "react";
import shallow from 'zustand/shallow'

import useStore from '../stores/fantacitorio'
import Loading from "../components/loading";
import Search from "../components/f-search";
import TweetList from "../components/f-tweet-list";
import Slideshow from "../components/f-slideshow";


const Fantacitorio: React.FC = () => {
  const { query, loading, fetch, scoreboard, teams } = useStore(s => ( {
    fetch: s.fetch,
    query: s.query,
    loading: s.loading,
    scoreboard: s.scoreboard,
    teams: s.teams
  }), shallow)

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
                <div className="row-start-2 lg:row-start-1 lg:row-span-2 col-span-1 flex flex-col xl:overflow-auto lg:flex-1 ">
                  <div className="lg:p-2 flex flex-col flex-initial box-border ">
                    <div className="flex items-center justify-center m-2">
                      <div className="flex flex-col">
                        <p className="border-b-2 border-gray-500"> <span className="text-orange-500">BEST AVERAGE</span>: {scoreboard.best_average.name} {scoreboard.best_average.surname}</p>
                        {/*<p className="border-b-2 border-gray-500"> <span className="text-orange-500">BEST CLIMBER</span>: {scoreboard.best_climber.name} {scoreboard.best_climber.surname}</p> */}
                        <p className="border-b-2 border-gray-500"> <span className="text-orange-500">BEST SINGLE SCORE</span>: {scoreboard.best_single_score.name} {scoreboard.best_single_score.name}</p>
                      </div>
                    
                    </div>
                    <div className="box-border flex items-center justify-evenly aspect-video lg:p-2 ">
                      <div className="flex flex-col">
                      {scoreboard.politicians.slice(0, 32).map((p, index) => (<p key={p.id}>{index+1}. <span className="text-orange-500">{p.points}</span> {p.name} {p.surname}</p>)
                      )}
                      </div>
                      <div className="flex flex-col">
                      {scoreboard.politicians.slice(32, 64).map((p, index) => (<p key={p.id}> {index+33}. <span className="text-orange-500">{p.points}</span> {p.name} {p.surname}</p>)
                      )}
                      </div>
                    </div>
                  </div>
                  <div className="p-2 flex items-center flex-col flex-initial flex-1 xl:overflow-none  aspect-video">
                    <Slideshow {...{teams}} />
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
