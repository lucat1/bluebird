package test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/kataras/muxie"
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

func staticHandler(buf []byte) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(buf)
	})
}

func CreateMultiServer(contentsMap map[string][]byte) *httptest.Server {
	mux := muxie.NewMux()
	for path, res := range contentsMap {
		mux.Handle(path, staticHandler(res))
	}
	return httptest.NewServer(mux)
}
