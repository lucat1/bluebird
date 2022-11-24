import React, { useEffect, useState } from 'react'
import { Chessboard } from 'react-chessboard'
import Countdown from 'react-countdown'
import { getLocalTimeZone } from '@internationalized/date'
import { useForm } from 'react-hook-form'
import { useElementSize } from 'usehooks-ts'
import { parseDateTime } from '@internationalized/date'
import useChess from '../stores/chess'
import { ChessState } from '../types'
import MoveList from '../components/move-list'

const Chess: React.FC = () => {
  const { state, fetch, turn, play, end, code } = useChess(s => s)
  const [authorized, setAuthorized] = useState(false)
  const [getRef, { width }] = useElementSize()
  console.log(state)
  const {
    handleSubmit,
    //formState: { errors },
    register,
  } = useForm<{ code: string }>({ reValidateMode: "onSubmit", mode: "onSubmit" });

  useEffect(() => {
    fetch();
  }, [])


  return (
    <div className='flex flex-1 flex-col md:flex-row'>

      {state == ChessState.IDLE && (
        <div className='flex flex-1 items-center justify-center '>
          <button
            onClick={_ => { play({ minutes: 1 }); setAuthorized(true) }}
            type="submit"
            className="text-white text-center hover:bg-sky-700  bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800"
          >
            Start!
          </button>
        </div>
      )}
      {state != ChessState.IDLE && !authorized && (
        <div className='flex flex-1 items-center justify-center '>
          <form onClick={handleSubmit(_ => setAuthorized(true))} className='flex'>
            <div>
              <label
                htmlFor="code"
                className="flex mb-2 text-sm font-medium text-gray-900 dark:text-gray-400"
              >
                Insert game code:
              </label>
              <div className='flex'>
                <input
                  id="code"
                  type="search"
                  className="flex p-4 pl-10 hover:border-gray-400   text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-sky-500 focus:border-sky-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-sky-500 dark:focus:border-sky-500"
                  placeholder="Codice..."
                  {...register("code", {
                    validate: value => value == code
                  })}
                />
                <button
                  type="submit"
                  className="text-white hover:bg-sky-700   bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800"
                >
                  Invia
                </button>

              </div>
            </div>
          </form>
        </div>
      )}
      {state != ChessState.IDLE && state != ChessState.LOST && authorized && (
        <>
          <div className="flex flex-1 flex-col p-9">
            <div ref={getRef}>
              <Chessboard
                boardWidth={width}
                arePiecesDraggable={turn}
                position="rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
              />
            </div>
            <div>
              {turn ? (
                <p>It's your turn: move a piece</p>) : (<>
                  Opponent's turn, waiting
                  <Countdown date={end!.toDate(getLocalTimeZone())} />
                </>)}
            </div>
          </div>
          <div className='flex max-w-5xl'>
            <div className='w-full'>
              <MoveList tweets={[{
                user: {
                  id: "Pino Daniele",
                  name: "Pino Daniele",
                  username: "Pino Daniele",
                  profile_image: "chupa"
                },
                id: "subeme",
                created_at: Date.now().toString(),
                text: "Ciao carissimi",
                date: parseDateTime('2022-02-03T09:15')
              }]}
              />
            </div>
          </div>
        </>
      )}
    </div>
  )
}

export default Chess
