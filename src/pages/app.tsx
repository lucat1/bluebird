import * as React from "react";
import { OverlayContainer } from "@react-aria/overlays";
import { ErrorBoundary } from 'react-error-boundary'

import useStore from '../stores/store'
import Error from '../components/error';
import Navbar from "../components/navbar";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import SearchPage from "./search-page";
import Ghigliottina from "./ghigliottina";

const App: React.FC = () => {
  const reset = useStore(s => s.reset)

  return (
    <OverlayContainer>
      <main className="w-screen h-screen flex flex-col overflow-auto lg:overflow-hidden dark:bg-gray-900 dark:text-gray-200 dark:text-white">
        <Navbar />
        <ErrorBoundary FallbackComponent={Error} onReset={reset}>
          <BrowserRouter>
            <Routes>
              <Route path="/" element={<SearchPage/>} />
              <Route path="/ghigliottina" element={<Ghigliottina/>} />
            </Routes>
          </BrowserRouter>
        </ErrorBoundary>
      </main>
    </OverlayContainer >
  );
}
export default App;
