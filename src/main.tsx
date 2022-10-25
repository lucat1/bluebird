import React from "react";
import ReactDOM from "react-dom/client";
import { Routes, Route, BrowserRouter } from "react-router-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import App from "./components/app";
import Search from "./pages/search";

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <BrowserRouter>
        <App>
          <Routes>
            <Route index element={<h1>home</h1>} />
            <Route path="search" element={<Search />} />
            <Route path="*" element={<h1>Not found</h1>} />
          </Routes>
        </App>
      </BrowserRouter>
    </QueryClientProvider>
  </React.StrictMode>
);
