import * as React from "react";

import useStore from "../stores/eredita"

const Classification: React.FC = () => {
const { ghigliottina } = useStore(s => ({ ghigliottina: s.ghigliottina, loadingGhigliottina: s.loadingGhigliottina }))


var first = new Date(ghigliottina!.podium.first.time);
var second = new Date(ghigliottina!.podium.second.time);
var third = new Date(ghigliottina!.podium.third.time);
return   (
  <div className="flex flex-col  py-3">
  <div className="flex felx-row items-center justify-between mb-4 px-2 border-b-2 border-gray-500">
    <div className="text-xl text-orange-500">1<sup>st</sup></div>
    <div className="flex space-x-4 px-2">
      <div>{ghigliottina!.podium.first.username}</div>
    </div>
    <span>{first.getHours()}:{first.getMinutes()}:{first.getSeconds()}</span>
  </div>
  <div className="flex felx-row items-center justify-between mb-4 px-2 border-b-2 border-gray-500">
  <div className="text-xl">2<sup>nd</sup></div>
    <div className="flex space-x-4 px-2">
      <div>{ghigliottina!.podium.second.username}</div>
    </div>
  <span >{second.getHours()}:{second.getMinutes()}:{second.getSeconds()}</span>
  </div>
  <div className="flex felx-row items-center justify-between  px-2 border-b-2 border-gray-500">
  <div className="text-xl">3<sup>rd</sup></div>
   <div className="flex space-x-4 px-2">
      <div>{ghigliottina!.podium.third.username}</div>
    </div>
  <span >{third.getHours()}:{third.getMinutes()}:{third.getSeconds()}</span>
  </div>
  
  </div>
  )
}

export default Classification
