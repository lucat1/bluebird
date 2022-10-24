package server

import (
	"fmt"
	"log"
	"net/http"

	"git.hjkl.gq/bluebird/bluebird/cache"
	"git.hjkl.gq/bluebird/bluebird/request"
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

func cors(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		f.ServeHTTP(w, r)
	})
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	var err error
	tweets := []request.Tweet{}
	if query != "" {
		searchType := r.URL.Query().Get("type")
		handler, has := searchHandlerMap[searchType]
		if !has {
			sendError(w, http.StatusInternalServerError, APIError{
				Message: "Unknown search type",
				Error:   fmt.Errorf("Search error"),
			})
			return
		}

		tweets, err = handler(query, 10)
		cache.AddTweets(tweets)
		if err != nil {
			sendError(w, http.StatusInternalServerError, APIError{
				Message: "Could not fetch tweets",
				Error:   err,
			})
			return
		}
	}
	sendJSON(w, 200, tweets)
}

func RunServer(host string) {
	mux := muxie.NewMux()

	mux.Use(cors)
	mux.HandleFunc("/api/search", searchHandler)
	mux.Handle("/*path", http.FileServer(http.Dir("dist")))

	log.Printf("Listening on %s\n", host)
	http.ListenAndServe(host, mux)
}
