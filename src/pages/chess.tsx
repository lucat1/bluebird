import React,{useEffect, useState} from 'react'
import { Chessboard } from 'react-chessboard'
import Countdown from 'react-countdown'
import { getLocalTimeZone } from '@internationalized/date'

import useChess from '../stores/chess'
import { ChessState } from '../types'

const Chess: React.FC = () => {
  const { state,fetch, turn, play, end, code } = useChess(s => s)
  const [authorized,setAuthorized]=useState(false)
  console.log(state)
  
  useEffect(()=>{
    fetch(); 
    },[])

  return (
    <div className='flex-1 grid grid-cols-[auto_min-content]'>
        
        {state == ChessState.IDLE && (
          <div className='col-span-2 flex items-center justify-center '>
                <button
                  onClick={_=>play({minutes: 1 })}
                  type="submit"
                  className="text-white text-center hover:bg-sky-700  bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800"
                >
                  Start!
                </button>
          </div>
        )}
        {state!= ChessState.IDLE && !authorized && (
          <form className='flex'>
          <div>
            <label
              htmlFor="code"
              className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-400"
            >
              Inserisci il codice della partita:
            </label>
            <input
              id="code"
              type="search"
              className="block p-4 pl-10 hover:border-gray-400   text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-sky-500 focus:border-sky-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-sky-500 dark:focus:border-sky-500"
              placeholder=" code..."
            />
              <button
                type="submit"
                className="text-white hover:bg-sky-700   bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800"
              >
                Invia
              </button>
              </div>
          </form> 
        )}
        {state != ChessState.IDLE && state != ChessState.LOST && authorized &&(
          <div className="flex p-9 ">
            <Chessboard
              arePiecesDraggable={turn}
              position="rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
            />
            {turn ? (
              <p>It's your turn: move a piece</p> ): (<>
              Opponent's turn, waiting
              <Countdown date={end!.toDate(getLocalTimeZone())} />
            </>)}
          </div>
        )}
    </div>
  )
}

export default Chess
