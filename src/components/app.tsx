import * as React from "react";
import { OverlayContainer } from "@react-aria/overlays";
import { ErrorBoundary } from 'react-error-boundary'

import Loading from "./loading";
import Navbar from "./navbar";
import Error from './error'
import { queryClient } from "../main";


const App: React.FC<React.PropsWithChildren<{}>> = ({ children }) => {
  
  return (
  <OverlayContainer>
    <main className="w-screen h-screen dark:bg-gray-900 dark:text-gray-800 dark:text-white overflow-auto">
    <Navbar></Navbar>

      <ErrorBoundary
        FallbackComponent={Error}
        onReset={() => queryClient.clear()}
      >
        <React.Suspense fallback={<Loading />}>{children}</React.Suspense>
      </ErrorBoundary>
    </main>
  </OverlayContainer>
);
}
export default App;
