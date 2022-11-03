import * as React from "react";
import { LatLng } from "leaflet";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import type { Tweet } from "../types";

const TweetMap: React.FC<{ tweets?: Tweet[] }> = ({ tweets }) => {
  const position = [51.505, -0.09];
  return (
    <details className="dark:bg-gray-900 mt-4 bg-white open:bg-orange-300 duration-300">
      <summary className="dark:bg-gray-900 dark:text-white text-center bg-inherit px-5 py-3 text-lg cursor-pointer">
        Open the map
      </summary>
      <div className="bg-white dark:bg-gray-900 px-5 py-3  text-sm font-light">
        <MapContainer
          className="m-auto shadow-lg"
          style={{ width: "80vw", height: "70vh" }}
          center={new LatLng(position[0], position[1])}
          zoom={13}
          scrollWheelZoom={true}
        >
          <TileLayer url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png" />
          {tweets?.map(
            (tweet, i) =>
              tweet.geo && (
                <Marker
                  key={i}
                  position={
                    new LatLng(
                      tweet.geo.coordinates[0],
                      tweet.geo.coordinates[1]
                    )
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
