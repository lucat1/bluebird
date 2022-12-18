import * as React from "react";
import { Carousel } from 'react-responsive-carousel';
import "react-responsive-carousel/lib/styles/carousel.min.css";
import { Team } from "../types";

const Slideshow: React.FC<{ teams: Team[] }> = (props) => {
  const [name, setName] = React.useState("")
  const filteredTeams = name ? props.teams.filter((team) => (team.username.includes(name))) : props.teams

  return (
    <div className="flex overflow-none flex-col w-full">
      <div className="border border-sky-500 overflow-none">
        {filteredTeams.length == 0 ?
          (<div className="overflow-none xl:max-w-xl xl:w-fit w-screen h-96 justify-center text-red-500 text-2xl text-center flex items-center justify-center">
            <span className="nline-block align-middle">Non esistono squadre create da quell'utente</span>
          </div>)
          : (<Carousel dynamicHeight={true} renderIndicator={false}>
            {filteredTeams.map((team) => (
              <div key={team.username} className="w-full aspect-video">
                <img src={team.picture_url} />
                <p className="legend" style={{ position: 'relative', top: '-2.25rem' }}>{team.username}</p>
              </div>)
            )}
          </Carousel>)}
      </div>
      <div className="flex flex-row border border-sky-500 p-1 justify-center">
        <input
          onChange={(e) => { setName(e.target.value) }}
          id="query"
          type="search"
          className="block p-4 hover:border-gray-400   text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-sky-500 focus:border-sky-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-sky-500 dark:focus:border-sky-500"
          placeholder="Search username"
        // {...register("query", { required: true })}
        />
      </div>
    </div>
  );
};

export default Slideshow;
