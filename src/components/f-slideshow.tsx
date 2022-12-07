import * as React from "react";
import "react-responsive-carousel/lib/styles/carousel.min.css"; // requires a loader
import { Carousel } from 'react-responsive-carousel';

const Slideshow: React.FC = () => {
  
  return (
    <div className="flex flex-col">
      <div className="border border-sky-500">
        <Carousel>
          <div>
            <img src="https://cdn.pixabay.com/photo/2016/11/06/05/36/lake-1802337__480.jpg" />
            <p className="legend">Legend 1</p>
          </div>
          <div>
            <img src="https://cdn.pixabay.com/photo/2016/11/06/05/36/lake-1802337__480.jpg" />
            <p className="legend">Legend 2</p>
          </div>
          <div>
            <img src="https://cdn.pixabay.com/photo/2016/11/06/05/36/lake-1802337__480.jpg" />
            <p className="legend">Legend 3</p>
          </div>
        </Carousel>
      </div>
      <div className="flex flex-row border border-sky-500 p-1 justify-between">
      <input
              id="query"
              type="search"
              className="block p-4 hover:border-gray-400   text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-sky-500 focus:border-sky-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-sky-500 dark:focus:border-sky-500"
              placeholder="Search"
              // {...register("query", { required: true })}
            />
            <div className="relative">
              <button
                type="submit"
                className="text-white absolute hover:bg-sky-700  right-2.5 bottom-2.5 bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800"
              >
                Search
              </button>
            </div>
      </div>
    </div>
  );
};

export default Slideshow;
