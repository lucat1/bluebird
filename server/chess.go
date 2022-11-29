package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"git.hjkl.gq/team14/team14/chess"
	"github.com/gorilla/websocket"
	"github.com/kataras/muxie"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(_ *http.Request) bool { return true },
	}
	connections     = []*websocket.Conn{}
	connectionsLock = sync.RWMutex{}
)

type OutgoingMessage[T any, D any] struct {
	Type    T      `json:"type"`
	Data    D      `json:"data"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}

type IncomingMessage[T any, D any] struct {
	Type T `json:"type"`
	Data D `json:"data"`
}

type ChessMessageType string

const (
	ChessMessageTypeMatch  ChessMessageType = "match"
	ChessMessageTypeStart                   = "start"
	ChessMessageTypeTweets                  = "tweets"
	ChessMessageTypeMove                    = "move"
)

func sendMessage[T any, D any](conn *websocket.Conn, msg OutgoingMessage[T, D]) {
	buf, err := json.Marshal(msg)
	if err != nil {
		log.Println("WARN: Could not serialize websocket message")
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, buf); err != nil {
		log.Println("WARN: Could not send websocket message")
		return
	}
}

func sendMatch() {
	var match *chess.SerializedMatch
	if chess.GetMatch() != nil {
		m := chess.GetMatch().Serialized()
		match = &m
	}

	connectionsLock.RLock()
	for _, conn := range connections {
		sendMessage(conn, OutgoingMessage[ChessMessageType, *chess.SerializedMatch]{
			Type: ChessMessageTypeMatch,
			Data: match,
		})
	}
	connectionsLock.RUnlock()
}

func onOpen(conn *websocket.Conn) {
	connectionsLock.Lock()
	connections = append(connections, conn)
	log.Printf("Opened a chess connection. Now at: %d", len(connections))
	connectionsLock.Unlock()
}

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func onClose(conn *websocket.Conn) {
	connectionsLock.Lock()
	i := -1
	for j, c := range connections {
		if c == conn {
			i = j
		}
	}
	if i != -1 {
		connections = remove(connections, i)
	}
	log.Printf("Closed a chess connection. Now at: %d", i)
	connectionsLock.Unlock()
	conn.Close()
}

func chessHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w.(*muxie.Writer).ResponseWriter, r, nil)
	if err != nil {
		log.Println("err", err)
		sendError(w, http.StatusBadRequest, APIError{
			Message: "Could not upgrade the connection to a WebSocket",
			Error:   err,
		})
		return
	}
	onOpen(conn)
	defer onClose(conn)

	for {
		_, raw, err := conn.ReadMessage()
		if err != nil {
			sendMessage(conn, OutgoingMessage[ChessMessageType, int]{
				Message: "Error reading message",
				Error:   err,
			})
			break
		}
		var msg IncomingMessage[ChessMessageType, string]
		if err := json.Unmarshal(raw, &msg); err != nil {
			sendMessage(conn, OutgoingMessage[ChessMessageType, int]{
				Message: "Invalid message format",
				Error:   err,
			})
			break
		}

		switch msg.Type {
		case ChessMessageTypeMatch:
			sendMatch()
			break

		case ChessMessageTypeStart:
			ms, err := strconv.Atoi(msg.Data)
			if err != nil {
				sendMessage(conn, OutgoingMessage[ChessMessageType, int]{
					Message: "Failed to parse duration",
					Error:   err,
				})
				break
			}

			match := chess.NewMatch(time.Millisecond * time.Duration(ms))
			chess.SetMatch(&match)
			chess.Store()
			sendMatch()
			break

		case ChessMessageTypeMove:
			if chess.GetMatch() == nil {
				msg := "Cannot move while a match hasn't been started"
				sendMessage(conn, OutgoingMessage[ChessMessageType, int]{
					Message: msg,
					Error:   errors.New(msg),
				})
				break
			}

			if err := chess.GetMatch().Move(msg.Data); err != nil {
				sendMessage(conn, OutgoingMessage[ChessMessageType, int]{
					Message: "Could not move",
					Error:   err,
				})
				break
			}

			chess.Store()
			sendMatch()
			break
		}
	}
}

func uploadBoard() {
}
