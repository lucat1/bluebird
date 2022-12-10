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
  timeout: NodeJS.Timer | null;

  game: string | null;
  gameover: "w" | "b" | "d" | "f" | null;
  tweets: Tweet[] | null;
  board: Chess | null;
  turn: Color | null;
}

export interface Actions {
  connect(): void;
  disconnect(): void;
  fetch(): void;
  getTweets(): void;
  _handler(e: MessageEvent<any>): void;
  _timeout(): void;
  play(turnDuration: TimeDuration): void;
  algebraic(from: Square, to: Square, piece: Pieces): string | null;
  move(mv: string): void;
  forfeit(): void;
}

const initialState: State = {
  connecting: true,
  connection: null,
  error: null,
  loading: true,
  timeout: null,

  code: "",
  end: null,

  gameover: null,
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
    try {
      const connection = new WebSocket(
        `${import.meta.env.DEV ? "ws://localhost:8080" : ""}/api/chess`
      );
      set({ connecting: true, connection });
      connection.onerror = (_) =>
        set({ ...get(), error: "Errore di connessione" });
      connection.onopen = (_) => {
        set({ ...get(), connecting: false, error: null });
        get().fetch();
      };
      connection.onmessage = (e) => get()._handler(e);
    } catch (_) {
      set({ ...get(), error: "Connection failed" });
    }
  },
  disconnect: async () => {
    get().connection?.close();
    set(initialState);
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

        // clear any previous timeout
        let { timeout, _timeout } = get();
        if (timeout) clearTimeout(timeout);

        const gameover = data.forfeited ? 'f' : board.isDraw()
          ? "d"
          : board.isCheckmate()
            ? board.turn() == "w"
              ? "b"
              : "w"
            : null;
        if (!gameover)
          timeout = setTimeout(
            _timeout,
            new Date(data.ends_at).getTime() - new Date().getTime()
          );

        set({
          ...data,
          loading: false,
          end,
          board,
          tweets: data.tweets?.map(convert),
          gameover,
          turn: board.turn(),
          timeout,
        });
        console.log("match", get());
        break;
    }
  },
  _timeout: () => {
    set({ loading: true, timeout: null });
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
    const { connection } = get();
    set({ ...initialState, connection, connecting: false, loading: true });
    connection?.send(
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
    const { connecting, connection, game, ...other } = get();
    if (connecting || !game || !connection) return;

    console.log("moving", move);
    connection?.send(
      JSON.stringify({
        type: ChessMessageType.Move,
        data: move,
      } as OutgoingMessage)
    );
    const tmp = new Chess(game);
    tmp.move(move);
    set({ ...other, connection, game: tmp.fen(), loading: true });
  },
  forfeit: () => {
    const { connecting, connection } = get();
    if (connecting || !connection) return null;

    console.log("surrendering");
    connection.send(JSON.stringify({ type: ChessMessageType.Forfeit, data: "" } as OutgoingMessage))
  }
}));

export default store;
