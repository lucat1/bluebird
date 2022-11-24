import create from 'zustand'
import { parseDateTime } from '@internationalized/date';
import type { TimeDuration, CalendarDateTime } from '@internationalized/date'
import { Chess, Color, Square } from 'chess.js'

import fetch, { withJSON } from '../fetch'
import type { Chess as RequestChess } from '../types'

export interface State {
  code: string | null
  timeout: number | null
  end: CalendarDateTime | null

  gameover: boolean
  game: string | null
  board: Chess | null
  turn: Color | null
}

export enum Player {
  WHITE, BLACK
}

export interface Actions {
  fetch(): Promise<void>
  play(turnDuration: TimeDuration): Promise<void>
  check(move: string): boolean
  move(move: string): Promise<void>
}

const initialState: State = {
  code: '',
  timeout: null,
  end: null,

  gameover: false,
  game: null,
  board: null,
  turn: null,
}


const store = create<State & Actions>((set, get) => ({
  ...initialState,
  fetch: async () => {
    const state = await fetch<RequestChess>('chess')
    const end = parseDateTime(state.ends_at.slice(0, -1))
    const board = new Chess(state.game)
    set({ ...state, end, board, gameover: board.isGameOver(), turn: board.turn() })

    if (!get().gameover) {
      console.trace('started timeout')
      if (get().timeout != null)
        clearTimeout(get().timeout!)

      set({
        ...get(), timeout: setTimeout(
          fetch,
          new Date(state.ends_at).getTime() - new Date().getTime()
        )
      })
    }
  },
  play: async (turnDuration: TimeDuration) => {
    await fetch<Chess>('chess/start', withJSON('POST', {
      duration: (turnDuration.milliseconds || 0) +
        ((turnDuration.seconds || 0) * 1000) +
        ((turnDuration.minutes || 0) * 1000 * 60) +
        ((turnDuration.hours || 0) * 1000 * 60 * 24)
    }))
    await get().fetch()
  },
  check: (move: string) => get().board!.moves().includes(move as any),
  move: async (move: string) => {
    await fetch<Chess>('chess/move', withJSON('POST', { move }))
    await get().fetch()
  }
}))

export default store
