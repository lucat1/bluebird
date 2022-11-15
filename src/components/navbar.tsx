import * as React from "react";
import shallow from "zustand/shallow";

import useStore, { QueryType } from '../store'

const Navbar: React.FC = () => {
  const [open, setOpen] = React.useState(false)
  const { query, fetch } = useStore(s => ({ query: s.query, fetch: s.fetch }), shallow)
  const search = (hashtag: string): React.MouseEventHandler<HTMLAnchorElement> => e => {
    e.preventDefault()
    fetch({ type: QueryType.Keyword, query: hashtag, timeRange: query.timeRange })
  }

  return (
    <nav className="flex flex-wrap justify-between bg-sky-800 border-gray-200 py-2 px-4 lg:px-8 dark:bg-gray-900">
      <a href="/" className="flex items-center">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          version="1.1"
          className="w-56 h-12"
          viewBox="0 0 1000 138"
        >
          <g transform="matrix(1,0,0,1,-0.6060606060606233,0.11062449270619368)">
            <svg
              viewBox="0 0 396 55"
              data-background-color="#003366"
              preserveAspectRatio="xMidYMid meet"
              height="138"
              width="1000"
              xmlns="http://www.w3.org/2000/svg"
            >
              <defs></defs>
              <g
                id="tight-bounds"
                transform="matrix(1,0,0,1,0.2400000000000091,-0.044089471730728746)"
              >
                <svg
                  viewBox="0 0 395.52 55.088178943461436"
                  height="55.088178943461436"
                  width="395.52"
                >
                  <g>
                    <svg
                      viewBox="0 0 492.4997962709097 68.5955625672284"
                      height="55.088178943461436"
                      width="395.52"
                    >
                      <g>
                        <rect
                          width="4.4761350383019005"
                          height="49.88316959483684"
                          x="407.2376309903191"
                          y="9.356196486195778"
                          fill="#ffffff"
                          opacity="1"
                          strokeWidth="0"
                          stroke="transparent"
                          fillOpacity="1"
                          className="rect-ei-0"
                          rx="1%"
                          id="ei-0"
                          data-palette-color="#ffffff"
                        ></rect>
                      </g>
                      <g transform="matrix(1,0,0,1,0,9.480283972750094)">
                        <svg
                          viewBox="0 0 395.52 49.63499462172821"
                          height="49.63499462172821"
                          width="395.52"
                        >
                          <g>
                            <svg
                              viewBox="0 0 395.52 49.63499462172821"
                              height="49.63499462172821"
                              width="395.52"
                            >
                              <g>
                                <svg
                                  viewBox="0 0 395.52 49.63499462172821"
                                  height="49.63499462172821"
                                  width="395.52"
                                >
                                  <g transform="matrix(1,0,0,1,0,0)">
                                    <svg
                                      width="395.52"
                                      viewBox="3.7 -34.4 280.18 35"
                                      height="49.63499462172821"
                                      data-palette-color="#ffffff"
                                    >
                                      <g
                                        className="undefined-text-0"
                                        data-fill-palette-color="primary"
                                        id="text-0"
                                      >
                                        <path
                                          d="M26.6-34.4Q29.2-34.4 31.38-33.33 33.55-32.25 34.83-30.3 36.1-28.35 36.1-25.95L36.1-25.95Q36.1-19.6 30.35-17.95L30.35-17.95 30.35-17.75Q36.9-16.25 36.9-9.15L36.9-9.15Q36.9-6.45 35.58-4.38 34.25-2.3 31.95-1.15 29.65 0 26.9 0L26.9 0 3.7 0 3.7-34.4 26.6-34.4ZM14.75-26.85L14.75-21 22.35-21Q23.45-21 24.18-21.78 24.9-22.55 24.9-23.7L24.9-23.7 24.9-24.2Q24.9-25.3 24.15-26.08 23.4-26.85 22.35-26.85L22.35-26.85 14.75-26.85ZM14.75-13.9L14.75-8 23.15-8Q24.25-8 24.98-8.78 25.7-9.55 25.7-10.7L25.7-10.7 25.7-11.2Q25.7-12.35 24.98-13.13 24.25-13.9 23.15-13.9L23.15-13.9 14.75-13.9ZM42.6 0L42.6-34.4 53.65-34.4 53.65-8.8 71.3-8.8 71.3 0 42.6 0ZM110.2-13.85Q110.2-6.8 105.75-3.1 101.3 0.6 93.09 0.6L93.09 0.6Q84.9 0.6 80.42-3.1 75.95-6.8 75.95-13.85L75.95-13.85 75.95-34.4 87-34.4 87-14Q87-11.1 88.55-9.38 90.09-7.65 93.05-7.65L93.05-7.65Q96-7.65 97.57-9.4 99.15-11.15 99.15-14L99.15-14 99.15-34.4 110.2-34.4 110.2-13.85ZM117.59 0L117.59-34.4 147.34-34.4 147.34-26.15 128.64-26.15 128.64-21.4 144.64-21.4 144.64-13.5 128.64-13.5 128.64-8.25 147.69-8.25 147.69 0 117.59 0ZM176.59-34.4Q179.19-34.4 181.37-33.33 183.54-32.25 184.81-30.3 186.09-28.35 186.09-25.95L186.09-25.95Q186.09-19.6 180.34-17.95L180.34-17.95 180.34-17.75Q186.89-16.25 186.89-9.15L186.89-9.15Q186.89-6.45 185.56-4.38 184.24-2.3 181.94-1.15 179.64 0 176.89 0L176.89 0 153.69 0 153.69-34.4 176.59-34.4ZM164.74-26.85L164.74-21 172.34-21Q173.44-21 174.17-21.78 174.89-22.55 174.89-23.7L174.89-23.7 174.89-24.2Q174.89-25.3 174.14-26.08 173.39-26.85 172.34-26.85L172.34-26.85 164.74-26.85ZM164.74-13.9L164.74-8 173.14-8Q174.24-8 174.97-8.78 175.69-9.55 175.69-10.7L175.69-10.7 175.69-11.2Q175.69-12.35 174.97-13.13 174.24-13.9 173.14-13.9L173.14-13.9 164.74-13.9ZM204.14 0L193.09 0 193.09-34.4 204.14-34.4 204.14 0ZM244.69-23.85Q244.69-20.7 242.99-18.15 241.29-15.6 238.04-14.4L238.04-14.4 245.69 0 233.29 0 227.19-12.55 223.09-12.55 223.09 0 212.04 0 212.04-34.4 233.09-34.4Q236.79-34.4 239.41-32.98 242.04-31.55 243.36-29.13 244.69-26.7 244.69-23.85L244.69-23.85ZM233.44-23.4Q233.44-24.75 232.54-25.65 231.64-26.55 230.34-26.55L230.34-26.55 223.09-26.55 223.09-20.2 230.34-20.2Q231.64-20.2 232.54-21.13 233.44-22.05 233.44-23.4L233.44-23.4ZM265.83-34.4Q283.88-34.4 283.88-17.2L283.88-17.2Q283.88 0 265.83 0L265.83 0 250.93 0 250.93-34.4 265.83-34.4ZM261.98-26.15L261.98-8.25 265.63-8.25Q272.58-8.25 272.58-15.7L272.58-15.7 272.58-18.7Q272.58-26.15 265.63-26.15L265.63-26.15 261.98-26.15Z"
                                          fill="#ffffff"
                                          data-fill-palette-color="primary"
                                        ></path>
                                      </g>
                                    </svg>
                                  </g>
                                </svg>
                              </g>
                            </svg>
                          </g>
                        </svg>
                      </g>
                      <g transform="matrix(1,0,0,1,423.43139701894006,0)">
                        <svg
                          viewBox="0 0 69.06839925196958 68.5955625672284"
                          height="68.5955625672284"
                          width="69.06839925196958"
                        >
                          <g>
                            <svg
                              xmlns="http://www.w3.org/2000/svg"
                              version="1.1"
                              x="0"
                              y="0"
                              viewBox="20.159999999999997 24.926320090634704 59.14 58.73513233492528"
                              enableBackground="new 0 0 100 100"
                              height="68.5955625672284"
                              width="69.06839925196958"
                              className="icon-icon-0"
                              id="icon-0"
                            >
                              <path
                                fill="#ff9966"
                                d="M66.547 36.555c-2.498 2.411-4.713 5.26-6.766 8.717-1.561 2.654-3.041 5.357-4.533 8.058l-1.117 2.039c-1.451 2.636-2.953 5.358-4.897 7.782-2.645 3.276-5.937 5.414-9.79 6.339L38.8 69.637c-1.05 0.214-2.138 0.329-3.263 0.379-1.171 1.993-2.462 3.913-3.961 5.702-2.138 2.558-4.605 4.712-7.697 6.05-0.059 0.027-0.099 0.063-0.216 0.133 0.203 0.081 0.354 0.143 0.502 0.196 3.86 1.395 7.852 1.771 11.913 1.465 3.104-0.241 6.104-0.888 8.922-2.256 3.918-1.893 6.742-4.873 8.667-8.74 1.494-2.996 2.443-6.173 3.385-9.365 1.209-4.059 2.381-8.135 3.715-12.159 1.299-3.912 3-7.657 5.297-11.105 2.846-4.25 6.516-7.343 11.539-8.623-0.748 0.047-1.482 0.127-2.195 0.274C72.246 32.232 69.346 33.859 66.547 36.555z"
                                data-fill-palette-color="accent"
                              ></path>
                              <path
                                fill="#ff9966"
                                d="M77.342 30.55c-1.387-0.159-2.758-0.108-4.145 0.176-1.811 0.374-3.6 1.11-5.463 2.261-2.545 1.566-4.914 3.547-7.25 6.061-1.834 1.976-3.607 4.012-5.381 6.044L53.9 46.477c-1.693 1.937-3.448 3.942-5.541 5.637-1.73 1.393-3.591 2.358-5.545 2.931-0.284 0.085-0.572 0.169-0.861 0.235-0.101 0.023-0.2 0.061-0.302 0.082-0.883 0.181-1.798 0.283-2.722 0.301-1.415 0.032-2.862-0.141-4.354-0.475-1.847 2.217-3.871 4.277-6.169 6.06-2.402 1.865-5.016 3.302-8.03 3.897-0.054 0.012-0.104 0.036-0.216 0.078 0.167 0.108 0.286 0.19 0.413 0.269 3.163 1.99 6.642 3.107 10.326 3.617 1.413 0.193 2.823 0.291 4.229 0.255 0.275-0.006 0.547-0.043 0.822-0.061 1.115-0.07 2.229-0.204 3.335-0.473 3.853-0.927 6.944-3.037 9.418-6.107 1.909-2.377 3.362-5.024 4.829-7.681 1.863-3.382 3.701-6.787 5.666-10.115 1.916-3.234 4.15-6.239 6.877-8.864 3.266-3.146 7.008-5.167 11.541-5.435C77.525 30.6 77.432 30.578 77.342 30.55z"
                                data-fill-palette-color="accent"
                              ></path>
                              <path
                                fill="#ff9966"
                                d="M78.934 29.812c-0.172-0.166-0.357-0.297-0.533-0.455-4.348-3.899-9.449-5.179-15.217-4.02-3.07 0.613-5.863 1.907-8.52 3.518-2.684 1.624-5.094 3.604-7.474 5.625-2.897 2.482-5.777 4.985-8.727 7.417-2.244 1.845-4.639 3.481-7.275 4.739-2.361 1.124-4.811 1.85-7.449 1.798-0.05-0.001-0.092 0.015-0.198 0.026 0.119 0.122 0.203 0.213 0.293 0.301 2.298 2.25 5.009 3.818 7.996 4.92 0.82 0.301 1.652 0.54 2.488 0.746 0.246 0.061 0.491 0.125 0.739 0.177 1.264 0.262 2.544 0.407 3.857 0.379 1.199-0.024 2.328-0.239 3.427-0.536 0.289-0.08 0.581-0.153 0.865-0.25 1.685-0.574 3.267-1.436 4.725-2.611 2.032-1.645 3.74-3.595 5.455-5.559 2.182-2.495 4.344-5.014 6.598-7.443 2.197-2.363 4.623-4.474 7.391-6.177 2.275-1.401 4.641-2.322 7.146-2.566 0.955-0.092 1.934-0.079 2.93 0.036 0.215 0.024 0.432 0.062 0.648 0.096 0.4 0.065 0.795 0.128 1.201 0.229C79.154 30.042 79.053 29.916 78.934 29.812zM70.387 29.444c-0.297 0.617-0.949 1.016-1.664 0.926-0.492-0.062-0.889-0.343-1.143-0.723-0.125-0.186-0.203-0.393-0.246-0.619-0.027-0.16-0.057-0.322-0.037-0.491 0.086-0.668 0.561-1.186 1.164-1.366 0.211-0.062 0.438-0.087 0.67-0.058 0.072 0.009 0.131 0.048 0.199 0.066 0.734 0.191 1.234 0.849 1.213 1.613 0 0.053 0.018 0.1 0.01 0.153C70.531 29.127 70.461 29.287 70.387 29.444z"
                                data-fill-palette-color="accent"
                              ></path>
                            </svg>
                            <g></g>
                          </g>
                        </svg>
                      </g>
                    </svg>
                  </g>
                </svg>
                <rect
                  width="395.52"
                  height="55.088178943461436"
                  fill="none"
                  stroke="none"
                  visibility="hidden"
                ></rect>
              </g>
            </svg>
          </g>
        </svg>
      </a>
      <button onClick={_ => setOpen(!open)} className="inline-flex justify-center items-center ml-3 text-white rounded-lg md:hidden hover:text-gray-300 focus:outline-none focus:ring-2 focus:ring-blue-300 dark:text-gray-400 dark:hover:text-white dark:focus:ring-gray-500" aria-controls="mobile-menu-2" aria-expanded="false">
        <svg className="w-6 h-6" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clipRule="evenodd"></path></svg>

      </button>

      <div className={`${open ? '' : 'hidden'} w-full w-full md:block md:w-auto`}>
        <ul className="flex flex-col p-4 mt-4 bg-gray-50 rounded-lg border border-gray-100 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0 md:bg-transparent dark:bg-gray-900 md:dark:bg-gray-900 dark:border-gray-600">
          <li>
            <a onClick={search('#ghigliottina')} className="cursor-pointer block py-2 pr-4 pl-3 text-gray-900 rounded md:text-white hover:bg-sky-700 hover:text-white md:hover:bg-transparent md:hover:text-orange-300 md:p-0 md:dark:hover:text-sky-600 dark:text-gray-400 dark:hover:bg-sky-600 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700" aria-current="page">Eredità</a>
          </li>
          <li>
            <a className="cursor-not-allowed disabled block py-2 pr-4 pl-3 text-gray-900 rounded md:text-white hover:bg-sky-700 hover:text-white md:hover:bg-transparent md:hover:text-orange-300 md:p-0 md:dark:hover:text-sky-600 dark:text-gray-400 dark:hover:bg-sky-600 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700">Reazione a Catena</a>
          </li>
          <li>
            <a className="cursor-not-allowed block py-2 pr-4 pl-3 text-gray-900 rounded md:text-white hover:bg-sky-700 hover:text-white md:hover:bg-transparent md:hover:text-orange-300 md:p-0 md:dark:hover:text-sky-600 dark:text-gray-400 dark:hover:bg-sky-600 dark:hover:text-white md:dark:hover:bg-transparent dark:border-gray-700">Scacchi</a>
          </li>
        </ul>

      </div>
    </nav>
  );
}

export default Navbar;
