import * as React from "react";
import { Map } from "leaflet";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";

import useStore from '../stores/store'
import type { Tweet } from "../types";

const findBounds = (points: [number, number][]): [[number, number], [number, number]] => {
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

const TweetMap: React.FC = () => {
  const tweets = useStore(s => s.tweets)
  const mappedTweets = React.useMemo(() => tweets.filter(t => t.geo).map((t): MappedTweet => {
    const geo = t.geo!
    return {
      ...t, coordinates:
        [geo.coordinates.length == 2 ? geo.coordinates[1] : (geo.coordinates[1] + geo.coordinates[3]) / 2,
        geo.coordinates.length == 2 ? geo.coordinates[0] : (geo.coordinates[0] + geo.coordinates[2]) / 2]
    }
  }) || [], [tweets]);

  const map = React.useRef<Map>();
  const center = React.useMemo((): [number, number] => {
    const [top, bottom] = findBounds(mappedTweets.map(t => t.coordinates))
    return [(top[0] + bottom[0]) / 2, (top[1] + bottom[1]) / 2]
  }, [mappedTweets])

  React.useEffect(() => {
    if (!map.current)
      return
    (map.current as any).invalidateSize()
  }, [map])

  return (
      <div className="flex justify-center bg-white dark:bg-gray-900 pb-4">
        <MapContainer
          className="shadow-lg rounded-md border border-gray-300 dark:border-gray-600 "
          style={{ width: "80vw", height: "70vh" }}
          center={center}
          zoom={3}
          scrollWheelZoom={true}
          ref={map as any}
        >
          <TileLayer url="https://tiles.stadiamaps.com/tiles/outdoors/{z}/{x}/{y}{r}.png" />
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
  );
};

export default TweetMap;
