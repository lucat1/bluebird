package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIError struct {
	Message string
	Error   error
}

type rawError struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func sendError(w http.ResponseWriter, code int, error APIError) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	buf, err := json.Marshal(rawError{
		Message: error.Message,
		Error:   error.Error.Error(),
	})
	if err != nil {
		panic(fmt.Sprintf("Could not send error: %v", err))
	}
	w.Write(buf)
}
