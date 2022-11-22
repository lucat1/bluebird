import create from 'zustand'
import { parseDateTime } from '@internationalized/date';
import type { TimeDuration, CalendarDateTime } from '@internationalized/date'

import fetch, { withJSON } from '../fetch'
import { Chess, ChessState } from '../types'

export interface State {
  state: ChessState
  timeout: number | null
  fen: string // FEN board representation
  turn: boolean
  end: CalendarDateTime | null
  code: string
}

export interface Actions {
  fetch(): Promise<void>
  play(turnDuration: TimeDuration): Promise<void>
  move(move: string): Promise<void>
}

const initialState: State = {
  state: ChessState.IDLE,
  timeout: null,
  fen: '',
  turn: false,
  end: null,
  code:''
}


const store = create<State & Actions>((set, get) => ({
  ...initialState,
  fetch: async () => {
    const state = await fetch<Chess>('chess')
    const end = parseDateTime(state.ends_at.slice(0, -1))
    set({ ...state, end })

    if (state.state == ChessState.WAITING && !state.turn) {
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
  move: async (move: string) => {
    await fetch<Chess>('chess/move', withJSON('POST', { move }))
    await get().fetch()
  }
}))

export default store
