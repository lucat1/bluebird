import * as React from 'react'
import { Chessboard } from 'react-chessboard'
import Countdown from 'react-countdown'
import { getLocalTimeZone } from '@internationalized/date'

import useChess from '../stores/chess'
import { ChessState } from '../types'

const Chess: React.FC = () => {
  const { state, turn, play, end } = useChess(s => s)
  console.log(state)


  return (
    <div className='flex'>
      {state == ChessState.IDLE && (<button onClick={_ => play({ minutes: 1 })}>Start!</button>)}
      {state != ChessState.IDLE && state != ChessState.LOST && (
        <div className="flex flex-1 items-center justify-center">
          {turn ? 'It\'s your turn: move a piece' : (<>
            Opponent's turn, waiting
            <Countdown date={end!.toDate(getLocalTimeZone())} />
          </>)}
          <Chessboard
            arePiecesDraggable={turn}
            position="rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR"
          />
        </div>
      )}
    </div>
  )
}

export default Chess
