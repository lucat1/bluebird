import * as React from "react";
import { LatLng } from "leaflet";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import type { Tweet } from "../types";

const TweetMap: React.FC<{ tweets?: Tweet[] }> = ({ tweets }) => {
  const position = [51.505, -0.09];
  return (
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
                new LatLng(tweet.geo.coordinates[0], tweet.geo.coordinates[1])
              }
            >
              <Popup>{tweet.text}</Popup>
            </Marker>
          )
      )}
    </MapContainer>
  );
};

export default TweetMap;
