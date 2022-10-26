import * as React from "react";
import { useQuery } from "@tanstack/react-query";
import {LatLng} from 'leaflet'
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import fetch from "../fetch";
import type { Tweet } from "../types";
 
export interface TweetProps {
  type: string;
  query: string;
}

const TweetList: React.FC<TweetProps> = ({ type, query }) => {
  const { data: tweets } = useQuery(
    ["search", type, query],
    () =>
      fetch<Tweet[]>(
        type && query ? `search?type=${type}&query=${query}` : `search`
      ),
    { suspense: true }
  );

  const position = [51.505, -0.09]
  return (
  <MapContainer className="m-auto shadow-lg" style={{ width: '80vw', height: '70vh' }} center={new LatLng(position[0], position[1])} zoom={13} scrollWheelZoom={true}>
    <TileLayer
      url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
    />
    <Marker position={new LatLng(position[0], position[1])}>
      <Popup>
        A pretty CSS3 popup. <br /> Easily customizable.
      </Popup>
    </Marker>
  </MapContainer>
  );
};

export default TweetList;
