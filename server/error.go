package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	message string
	error   error
}

func sendError(w http.ResponseWriter, code int, error Error) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	buf, err := json.Marshal(error)
	if err != nil {
		panic(fmt.Sprintf("Could not send error: %e", err))
	}
	w.Write(buf)
}
