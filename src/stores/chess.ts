import create from "zustand";
import { parseDateTime } from "@internationalized/date";
import type { TimeDuration, CalendarDateTime } from "@internationalized/date";
import { Chess, Color, Square } from "chess.js";

import fetch, { withJSON } from "../fetch";
import {
  Chess as RequestChess,
  IncomingMessage,
  ChessMessageType,
  OutgoingMessage,
  Match as ChessMatch,
} from "../types";

export interface State {
  connection: WebSocket | null;
  connecting: boolean;
  error: string | null;
  loading: boolean;

  code: string | null;
  end: CalendarDateTime | null;

  gameover: boolean;
  game: string | null;
  board: Chess | null;
  turn: Color | null;
}

export enum Player {
  WHITE,
  BLACK,
}

export interface Actions {
  connect(): void;
  fetch(): void;
  _handler(e: MessageEvent<any>): void;
  play(turnDuration: TimeDuration): void;
  check(move: string): boolean;
  move(move: string): void;
}

const initialState: State = {
  connecting: true,
  connection: null,
  error: null,
  loading: true,
  code: "",
  end: null,

  gameover: false,
  game: null,
  board: null,
  turn: null,
};

const store = create<State & Actions>((set, get) => ({
  ...initialState,
  connect: async () => {
    set({ connecting: true });
    await new Promise((res) => setTimeout(res, 1000))
    const connection = new WebSocket(
      `${import.meta.env.DEV ? "ws://localhost:8080" : ""}/api/chess`
    );
    set({ connecting: true, connection });
    connection.onerror = (_) => set({ ...get(), error: "Connection failed" });
    connection.onopen = (_) => {
      set({ ...get(), connecting: false, error: null });
      get().fetch();
    };
    connection.onmessage = (e) => get()._handler(e);
  },
  _handler: (e: MessageEvent<any>) => {
    const msg = JSON.parse(e.data) as IncomingMessage<any>;
    console.log(msg);
    switch (msg.type) {
      case ChessMessageType.Match:
        const { data } = msg as IncomingMessage<ChessMatch>;
        if (data == null) {
          set({ ...get(), loading: false });
          return;
        }
        const end = parseDateTime(data.ends_at.slice(0, -1));
        const board = new Chess(data.game);
        set({
          ...data,
          loading: false,
          end,
          board,
          gameover: board.isGameOver(),
          turn: board.turn(),
        });
        break;
    }
  },
  fetch: () => {
    set({ ...get(), loading: true });
    get().connection?.send(
      JSON.stringify({
        type: ChessMessageType.Match,
        data: "",
      } as OutgoingMessage)
    );
  },
  play: (turnDuration: TimeDuration) => {
    set({ ...get(), loading: true });
    get().connection?.send(
      JSON.stringify({
        type: ChessMessageType.Start,
        data: (
          (turnDuration.milliseconds || 0) +
          (turnDuration.seconds || 0) * 1000 +
          (turnDuration.minutes || 0) * 1000 * 60 +
          (turnDuration.hours || 0) * 1000 * 60 * 60
        ).toString(),
      } as OutgoingMessage)
    );
  },
  check: (move: string) =>
    get()
      .board!.moves()
      .includes(move as any),
  move: (move: string) => {
    set({ ...get(), loading: true });
    get().connection?.send(
      JSON.stringify({
        type: ChessMessageType.Move,
        data: move,
      } as OutgoingMessage)
    );
  },
}));

export default store;
