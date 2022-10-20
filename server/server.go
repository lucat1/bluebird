package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"git.hjkl.gq/bluebird/bluebird/request"
	"github.com/aymerick/raymond"
	"github.com/kataras/muxie"
)

type indexPayload struct {
	Query  string
	Tweets []request.Tweet
}

func keywordHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf8")

	tplByte, err := ioutil.ReadFile("views/index.tpl")
	if err != nil {
		panic(err)
	}
	tpl := string(tplByte)

	query := r.URL.Query().Get("query")
	tweets := []request.Tweet{}
	if query != "" {
		tweets, err = request.TweetsByKeyword(query, 10)
		if err != nil {
			panic(err)
		}
	}

	result, err := raymond.Render(tpl, indexPayload{Query: query, Tweets: tweets})
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, result)
}

func RunServer(host string) {
	mux := muxie.NewMux()

	mux.HandleFunc("/search/keyword", keywordHandler)

	log.Printf("Listening on %s\n", host)
	http.ListenAndServe(host, mux)
}
