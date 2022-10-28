package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"git.hjkl.gq/bluebird/bluebird/cache"
	"git.hjkl.gq/bluebird/bluebird/request"
	"github.com/kataras/muxie"
)

const nOfAPITweets uint = 30

type indexPayload struct {
	Query  string
	Tweets []request.Tweet
}

type Fetcher func(string, uint) (tweets []request.Tweet, err error)

var searchHandlerMap = map[string]Fetcher{
	"keyword": request.TweetsByKeyword,
	"user":    request.TweetsByUser,
}

func cors(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		f.ServeHTTP(w, r)
	})
}

type SearchResponse struct {
	Tweets []request.Tweet `json:"tweets"`
	Cached uint            `json:"cached"`
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	tweets := []request.Tweet{}
	cached := uint(0)
	if query != "" {
		var err error
		searchType := r.URL.Query().Get("type")
		amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
		if err != nil {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Invalid amount query",
				Error:   err,
			})
			return
		}
		handler, has := searchHandlerMap[searchType]
		if !has {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Unknown search type",
				Error:   fmt.Errorf("Search error"),
			})
			return
		}

		tweets, err = handler(query, nOfAPITweets)
		if err != nil {
			sendError(w, http.StatusInternalServerError, APIError{
				Message: "Could not fetch tweets",
				Error:   err,
			})
			return
		}
		cached = uint(amount - len(tweets))
		if err = cache.InsertTweets(tweets); err != nil {
			sendError(w, http.StatusInternalServerError, APIError{
				Message: "Could not store gathered tweets in cache",
				Error:   err,
			})
			return
		}

		tweets, err = cache.TweetsByKeyword(query, uint(amount))
		if err != nil {
			sendError(w, http.StatusInternalServerError, APIError{
				Message: "Could not fetch tweets from cache",
				Error:   err,
			})
			return
		}
	}
	sendJSON(w, 200, SearchResponse{
		Tweets: tweets,
		Cached: cached,
	})
}

func RunServer(host string) {
	mux := muxie.NewMux()

	mux.Use(cors)
	mux.HandleFunc("/api/search", searchHandler)
	mux.Handle("/*path", http.FileServer(http.Dir("dist")))

	log.Printf("Listening on %s\n", host)
	http.ListenAndServe(host, mux)
}
