package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"git.hjkl.gq/bluebird/bluebird/cache"
	"git.hjkl.gq/bluebird/bluebird/request"
	"github.com/kataras/muxie"
)

const (
	nOfAPITweets   uint = 30
	nOfDaysAllowed      = 7
)

type indexPayload struct {
	Query  string
	Tweets []request.Tweet
}

type Fetcher func(string, uint, string, string) (tweets []request.Tweet, err error)

var twitterHandlerMap = map[string]Fetcher{
	"keyword": request.TweetsByKeyword,
	"user":    request.TweetsByUser,
}

var searchHandlerMap = map[string]Fetcher{
	"keyword": cache.TweetsByKeyword,
	"user":    cache.TweetsByUser,
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

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dist/index.html")
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
		handler1, has := twitterHandlerMap[searchType]
		if !has {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Unknown search type",
				Error:   fmt.Errorf("Search error"),
			})
			return
		}
		handler2, has := searchHandlerMap[searchType]
		if !has {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Unknown search type (2)",
				Error:   fmt.Errorf("Search error"),
			})
			return
		}

		rawStartTime := r.URL.Query().Get("startTime")
		rawEndTime := r.URL.Query().Get("endTime")
		startTime, err := time.Parse(time.RFC3339, rawStartTime)
		if err != nil {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Invalid startTime",
				Error:   err,
			})
			return
		}
		endTime, err := time.Parse(time.RFC3339, rawEndTime)
		if err != nil {
			sendError(w, http.StatusBadRequest, APIError{
				Message: "Invalid endTime",
				Error:   err,
			})
			return
		}
		maxTime := time.Now().Add(time.Hour * 24 * time.Duration(-nOfDaysAllowed))
		// fixing the startTime is only worth if the user is acutally interested in
		// the last N_OF_DAYS_ALLOWED days of activity
		if time.Since(startTime).Hours()/24 > nOfDaysAllowed && endTime.After(maxTime) {
			startTime = maxTime
		}

		tweets, err = handler1(query, nOfAPITweets, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339))
		if err != nil {
			log.Printf("Error while querying Twitter: %v", err)
		}
		cached = uint(amount - len(tweets))
		if len(tweets) > 0 {
			if err = cache.InsertTweets(tweets); err != nil {
				sendError(w, http.StatusInternalServerError, APIError{
					Message: "Could not store gathered tweets in cache",
					Error:   err,
				})
				return
			}
		}

		tweets, err = handler2(query, uint(amount), rawStartTime, rawEndTime)
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

func RunServer(host string) error {
	mux := muxie.NewMux()

	mux.Use(cors)
	mux.HandleFunc("/api/search", searchHandler)
	mux.Handle("/assets/*path", http.FileServer(http.Dir("dist")))
	mux.HandleFunc("/*path", serveIndex)

	log.Printf("Listening on %s\n", host)
	return http.ListenAndServe(host, mux)
}
