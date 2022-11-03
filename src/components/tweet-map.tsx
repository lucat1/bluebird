import * as React from "react";
import { LatLng } from "leaflet";
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
  const mappedTweets = tweets?.filter(t => t.geo).map((t): MappedTweet => {
    const geo = t.geo!
    return {
      ...t, coordinates:
        [geo.coordinates.length == 2 ? geo.coordinates[1] : (geo.coordinates[1] + geo.coordinates[3]) / 2,
        geo.coordinates.length == 2 ? geo.coordinates[0] : (geo.coordinates[0] + geo.coordinates[2]) / 2]
    }
  }) || [];
  const [top, bottom] = findBounds(mappedTweets.map(t => t.coordinates))
  const center: [number, number] = [(top[0] + bottom[0]) / 2, (top[1] + bottom[1]) / 2]
  return (
    <details className="dark:bg-gray-900 mt-4 bg-white open:bg-orange-300 duration-300">
      <summary className="dark:bg-gray-900 dark:text-white text-center bg-inherit px-5 py-3 text-lg cursor-pointer">
        Open the map
      </summary>
      <div className="bg-white dark:bg-gray-900 px-5 py-3  text-sm font-light">
        <MapContainer
          className="m-auto shadow-lg"
          style={{ width: "80vw", height: "70vh" }}
          center={center}
          zoom={13}
          scrollWheelZoom={true}
        >
          <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
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
    </details>
  );
};

export default TweetMap;
