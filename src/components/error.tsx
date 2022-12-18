import * as React from "react";
import { FallbackProps } from "react-error-boundary";

const Error: React.FC<FallbackProps> = ({ resetErrorBoundary }) => {
  return (
    <main className="w-screen h-screen flex flex-col items-center justify-center">
      <h1 className="my-4 text-2xl font-bold">An error occoured</h1>
      <button
        onClick={(_) => resetErrorBoundary()}
        className="text-white bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800"
      >
        Reset
      </button>
    </main>
  );
};

export default Error;
