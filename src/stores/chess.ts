import create from "zustand";
import { parseDateTime } from "@internationalized/date";
import type { TimeDuration, CalendarDateTime } from "@internationalized/date";
import { Pieces, Square } from "react-chessboard";
import { Chess, Color, PAWN } from "chess.js";

import { convert } from "./store";
import {
  IncomingMessage,
  ChessMessageType,
  OutgoingMessage,
  Match as ChessMatch,
  Tweet,
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
  tweets: Tweet[] | null;
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
  getTweets(): void;
  _handler(e: MessageEvent<any>): void;
  play(turnDuration: TimeDuration): void;
  algebraic(from: Square, to: Square, piece: Pieces): string | null;
  move(mv: string): void;
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
  tweets: null,
  board: null,
  turn: null,
};

const algebraic = (
  board: Chess,
  from: Square,
  to: Square,
  piece: Pieces
): string[] => {
  const capturing = board.get(to);
  const p = piece.charAt(1);
  const base = p.toLowerCase() != PAWN ? p : "";

  if (!capturing) return [base + to];
  else return [base + from.charAt(0) + "x" + to, base + "x" + to];
};

const store = create<State & Actions>((set, get) => ({
  ...initialState,
  connect: async () => {
    set({ connecting: true });
    await new Promise((res) => setTimeout(res, 1000));
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
          tweets: data.tweets?.map(convert),
          gameover: board.isGameOver(),
          turn: board.turn(),
        });
        console.log("match", get());
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
  getTweets: () => {
    set({ ...get(), loading: true });
    get().connection?.send(
      JSON.stringify({
        type: ChessMessageType.Tweets,
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
  algebraic: (from, to, piece) => {
    const { board } = get();
    if (!board) return null;

    const mvs = algebraic(board, from, to, piece);
    const nb = new Chess(board.fen());
    for (const mv of mvs) {
      console.info("testing move", mv);
      if (nb.move(mv) != null) return mv;
    }
    return null;
  },
  move: (move) => {
    console.log("moving", move);
    get().connection?.send(
      JSON.stringify({
        type: ChessMessageType.Move,
        data: move,
      } as OutgoingMessage)
    );
  },
}));

export default store;
