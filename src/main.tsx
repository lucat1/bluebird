import ReactDOM from "react-dom";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

import App from "./components/app";
import Search from "./search";

const queryClient = new QueryClient();

ReactDOM.render(
  <QueryClientProvider client={queryClient}>
    <App>
      <Search />
    </App>
  </QueryClientProvider>,
  document.getElementById("root")!
);
