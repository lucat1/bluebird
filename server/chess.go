package server

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"git.hjkl.gq/team14/team14/chess"
)

type StartMatchRequest struct {
	Duration int64 `json:"duration"`
}

func startMatch(w http.ResponseWriter, r *http.Request) {
	if chess.GetMatch() != nil {
		msg := "A match is already in progress"
		sendError(w, http.StatusBadRequest, APIError{
			Message: msg,
			Error:   errors.New(msg),
		})
		return
	}

	if r.Method == http.MethodPost {
		var dur StartMatchRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Could not read request body",
				Error:   err,
			})
			return
		}
		if err = json.Unmarshal(body, &dur); err != nil {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Could parse request body",
				Error:   err,
			})
			return
		}

		match := chess.NewMatch(time.Millisecond * time.Duration(dur.Duration))
		chess.SetMatch(&match)
	}
	sendJSON(w, http.StatusOK, nil)
}

type MoveRequest struct {
	Move string `json:"move"`
}

func move(w http.ResponseWriter, r *http.Request) {
	if chess.GetMatch() == nil {
		msg := "A match hasn't been started yet"
		sendError(w, http.StatusBadRequest, APIError{
			Message: msg,
			Error:   errors.New(msg),
		})
		return
	}

	if r.Method == http.MethodPost {
		var mv MoveRequest
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Could not read request body",
				Error:   err,
			})
			return
		}
		if err = json.Unmarshal(body, &mv); err != nil {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Could parse request body",
				Error:   err,
			})
			return
		}

		if err := chess.GetMatch().Move(mv.Move); err != nil {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Could not move",
				Error:   err,
			})
			return
		}
		chess.Store()
		sendJSON(w, http.StatusOK, chess.GetMatch().Serialized())
		return
	}
	sendJSON(w, http.StatusOK, nil)
}

func getMatch(w http.ResponseWriter, r *http.Request) {
	match := chess.GetMatch()
	if match != nil {
		sendJSON(w, http.StatusOK, match.Serialized())
	} else {
		sendJSON(w, http.StatusOK, nil)
	}
}

func uploadBoard() {
}
