package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"git.hjkl.gq/bluebird/bluebird/request"
	"github.com/aymerick/raymond"
	"github.com/kataras/muxie"
)

type indexPayload struct {
	Tweets []request.Tweet
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf8")

	tplByte, err := ioutil.ReadFile("views/index.tpl")
	if err != nil {
		panic(err)
	}
	tpl := string(tplByte)

	escaped, err := url.QueryUnescape(muxie.GetParam(w, "hashtag"))
	if err != nil {
		panic(err)
	}
	tweets, err := request.TweetsByKeyword("#"+escaped, 10)
	if err != nil {
		panic(err)
	}

	result, err := raymond.Render(tpl, indexPayload{Tweets: tweets})
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, result)
}

func RunServer(host string) {
	mux := muxie.NewMux()

	mux.HandleFunc("/search/hashtag/:hashtag", indexHandler)

	log.Printf("Listening on %s\n", host)
	http.ListenAndServe(host, mux)
}
