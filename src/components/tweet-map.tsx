import * as React from "react";
import { useQuery } from "@tanstack/react-query";
import { LatLng } from "leaflet";
import { MapContainer, TileLayer, Marker, Popup } from "react-leaflet";
import fetch from "../fetch";
import type { Search } from "../types";

export interface TweetProps {
  type: string;
  query: string;
}

const TweetList: React.FC<TweetProps> = ({ type, query }) => {
  const { data: tweets } = useQuery(
    ["search", type, query],
    () =>
      fetch<Search>(
        type && query ? `search?type=${type}&query=${query}&amount=50` : `search`
      ),
    { suspense: true }
  );

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
      {tweets?.tweets.map(
        (tweet) =>
          tweet.geo && (
            <Marker
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

export default TweetList;
