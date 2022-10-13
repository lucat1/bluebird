package api

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Cotent-Type", "application/json")
	encoded, err := json.Marshal(Message{Message: "Hello World!"})
	if err != nil {
		w.Write([]byte("Internal Server Error"))
		w.WriteHeader(500)
	}
	w.Write(encoded)
	w.WriteHeader(200)
}
