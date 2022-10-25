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

var searchHandlerMap = map[string]func(string, uint) (tweets []request.Tweet, err error){
	"keyword": request.TweetsByKeyword,
	"user":    request.TweetsByUser,
}

func searchHandler(w http.ResponseWriter, r *http.Request) {

	tplByte, err := ioutil.ReadFile("views/index.tpl")
	if err != nil {
		sendError(w, http.StatusInternalServerError, APIError{
			Message: "Could not read template",
			Error:   err,
		})
		return
	}
	tpl := string(tplByte)

	query := r.URL.Query().Get("query")
	tweets := []request.Tweet{}
	if query != "" {
		search_type := r.URL.Query().Get("type")
		handler, has := searchHandlerMap[search_type]
		if !has {
			sendError(w, http.StatusInternalServerError, APIError{
				Message: "Unknown search type",
				Error:   fmt.Errorf("Search error"),
			})
			return
		}

		tweets, err = handler(query, 10)
		if err != nil {
			sendError(w, http.StatusInternalServerError, APIError{
				Message: "Could not fetch tweets",
				Error:   err,
			})
			return
		}
	}

	result, err := raymond.Render(tpl, indexPayload{Query: query, Tweets: tweets})
	if err != nil {
		sendError(w, http.StatusInternalServerError, APIError{
			Message: "Could not render tweets",
			Error:   err,
		})
		return
	}

	w.Header().Set("Content-Type", "text/html;charset=utf8")
	fmt.Fprintf(w, result)
}

func RunServer(host string) {
	mux := muxie.NewMux()

	mux.HandleFunc("/search", searchHandler)

	log.Printf("Listening on %s\n", host)
	http.ListenAndServe(host, mux)
}
