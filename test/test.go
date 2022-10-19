package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func ReadFile(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return bytes
}

func ReadJSON(filename string, dest interface{}) {
	bytes := ReadFile(filename)
	if err := json.Unmarshal(bytes, dest); err != nil {
		panic(err)
	}
}

func CreateServer(contents []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(contents)
	}))
}
