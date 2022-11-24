package server

import (
	"encoding/json"
	"net/http"
)

func sendJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	buf, err := json.Marshal(data)
	if err != nil {
		sendError(w, 500, APIError{
			Message: "Could not marshal json response",
			Error:   err,
		})
	}
	w.Write(buf)
}
