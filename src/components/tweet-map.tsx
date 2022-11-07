import * as React from "react";
import { LatLng, Map } from "leaflet";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import type { Geo, Tweet } from "../types";

function findBounds(points: [number, number][]): [[number, number], [number, number]] {
  if (points.length <= 0) return [[0, 0], [0, 0]];
  const lo = points[0].slice() as [number, number]
  const hi = points[0].slice() as [number, number]
  for (const p of points) {
    lo[0] = Math.min(lo[0], p[0])
    hi[0] = Math.max(hi[0], p[0])
    lo[1] = Math.min(lo[1], p[1])
    hi[1] = Math.max(hi[1], p[1])
  }
  return [lo, hi]
}

interface MappedTweet extends Tweet {
  coordinates: [number, number]
}

const TweetMap: React.FC<{ tweets?: Tweet[] }> = ({ tweets }) => {
  const mappedTweets = React.useMemo(() => tweets?.filter(t => t.geo).map((t): MappedTweet => {
    const geo = t.geo!
    return {
      ...t, coordinates:
        [geo.coordinates.length == 2 ? geo.coordinates[1] : (geo.coordinates[1] + geo.coordinates[3]) / 2,
        geo.coordinates.length == 2 ? geo.coordinates[0] : (geo.coordinates[0] + geo.coordinates[2]) / 2]
    }
  }) || [], [tweets]);
  const [show, setShow] = React.useState(false);
  const map = React.useRef<Map>();
  const [top, bottom] = findBounds(mappedTweets.map(t => t.coordinates))
  const center: [number, number] = [(top[0] + bottom[0]) / 2, (top[1] + bottom[1]) / 2]
  React.useEffect(() => {
    if (!map.current)
      return
    (map.current as any).invalidateSize()
  }, [map, show])
  return (
    <div className="flex flex-col items-center justify-center">
      <button onClick={_ => setShow(!show)} className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-blue-600 dark:hover:bg-blue-700 focus:outline-none dark:focus:ring-blue-800 w-64">Show map</button>
      <div className={`flex justify-center bg-white dark:bg-gray-900 ${show ? '' : 'hidden'}`}>
        <MapContainer
          className="shadow-lg rounded-md border border-gray-300 dark:border-gray-600"
          style={{ width: "80vw", height: "70vh" }}
          center={center}
          zoom={3}
          scrollWheelZoom={true}
          ref={map}
        >
          <TileLayer url="https://tiles.stadiamaps.com/tiles/outdoors/{z}/{x}/{y}{r}.png" />
          {/*<TileLayer url="https://mts1.google.com/vt/lyrs=m@186112443&hl=x-local&src=app&x={x}&y={y}&z={z}" />*/}
          {mappedTweets.map(
            (tweet, i) =>
              tweet.geo && (
                <Marker
                  key={i}
                  position={
                    tweet.coordinates
                  }
                >
                  <Popup>{tweet.text}</Popup>
                </Marker>
              )
          )}
        </MapContainer>
      </div>
    </div>
  );
};

export default TweetMap;
