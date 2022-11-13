package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"git.hjkl.gq/team14/team14/cache"
	"git.hjkl.gq/team14/team14/request"
	"github.com/kataras/muxie"
)

const (
	nOfAPITweets   uint = 100
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

type SentimentResponse struct {
	Sentiments request.Sentiments `json:"sentiments"`
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
		maxTime := time.Now().Add((time.Hour * (24*time.Duration(-nOfDaysAllowed) + 1)) + (time.Minute * time.Duration(59)))
		// fixing the startTime is only worth if the user is acutally interested in
		// the last N_OF_DAYS_ALLOWED days of activity
		if time.Since(startTime).Hours()/24 > nOfDaysAllowed && endTime.After(maxTime) {
			startTime = maxTime
		}
		if time.Now().Sub(endTime).Seconds() < 10 {
			endTime = time.Now()
			endTime.Add(time.Second * time.Duration(-10))
		}
		endTime.Add(time.Second * time.Duration(-10))

		log.Printf("Querying: \"%s\" %d %v %v", query, nOfAPITweets, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339))
		tweets, err = handler1(query, nOfAPITweets, startTime.Format(time.RFC3339), endTime.Format(time.RFC3339))
		if err != nil {
			log.Printf("Error while querying Twitter: %v", err)
		}
		cached = uint(len(tweets))
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
		cached = uint(len(tweets) - int(cached))
	}
	sendJSON(w, 200, SearchResponse{
		Tweets: tweets,
		Cached: cached,
	})
}

func sentimentHandler(w http.ResponseWriter, r *http.Request) {
	tweetID := r.URL.Query().Get("id")
	sentiments := request.Sentiments{}
	if tweetID != "" {
		tweet, err := cache.TweetByID(tweetID)
		if err != nil {
			sendError(w, http.StatusInternalServerError, APIError{
				Message: "Could not fetch tweet from cache",
				Error:   err,
			})
			return
		}
		if tweet.Sentiments != nil {
			sentiments = *tweet.Sentiments
		} else {
			sentiments, err = request.SentimentsFromTweet(tweet)
			if err != nil {
				sendError(w, http.StatusInternalServerError, APIError{
					Message: "Could not estimate sentiments",
					Error:   err,
				})
				return
			}
			tweet.Sentiments = &sentiments
			if err = cache.InsertTweets([]request.Tweet{tweet}); err != nil {
				sendError(w, http.StatusInternalServerError, APIError{
					Message: "Could not store gathered tweets in cache",
					Error:   err,
				})
				return
			}
		}
	}
	sendJSON(w, 200, SentimentResponse{
		Sentiments: sentiments,
	})
}

func RunServer(host string) error {
	mux := muxie.NewMux()

	mux.Use(cors)
	mux.HandleFunc("/api/search", searchHandler)
	mux.HandleFunc("/api/sentiment", sentimentHandler)
	mux.Handle("/assets/*path", http.FileServer(http.Dir("dist")))
	mux.HandleFunc("/*path", serveIndex)

	log.Printf("Listening on %s\n", host)
	return http.ListenAndServe(host, mux)
}
