import * as React from "react";
import { OverlayContainer } from "@react-aria/overlays";
import { ErrorBoundary } from 'react-error-boundary'

import Loading from "./loading";
import Navbar from "./navbar";
import Error from './error';
import TermCloud from "./term-cloud";
import Search from "../search";
import { queryClient } from "../main";


const App: React.FC<React.PropsWithChildren<{}>> = ({ children }) => {
  return (
  <OverlayContainer>
    <main className="w-screen h-screen overflow-auto lg:overflow-none flex flex-col dark:bg-gray-900 dark:text-gray-200 dark:text-white">
      <Navbar></Navbar>
      <div className="lg:flex-1 grid grid-rows-none lg:grid-cols-[2fr,1fr] lg:grid-rows-[min-content,auto]">
        <div className=" lg:grid lg:grid-row-[min-content,auto]  row-start-2 lg:row-span-2 lg:row-start-1">
          <div className="lg:grid lg:grid-cols-[1fr,2fr] lg:row-start-1">
            <div className="lg:col-start-1 p-2"> torta</div>
            <div className="lg:col-start-2 p-2"> barre</div>
          </div>
          <div className="lg:row-start-2 p-2"> TermCloud</div>
          <div className="lg:row-start-3 p-2"> bottone mappa</div>
        </div>
        <div className=" justify-center flex lg:col-start-2 p-2">
          <Search></Search>
        </div>
        <div className=" lg:row-start-2 lg:col-start-2">tweet list</div>
      </div>

    </main>
  </OverlayContainer>
);
}
export default App;
