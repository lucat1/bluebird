package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"git.hjkl.gq/team14/team14/cache"
	"git.hjkl.gq/team14/team14/request"
)

type Fetcher func(string, uint, *time.Time, *time.Time) (tweets []request.Tweet, err error)

var twitterHandlerMap = map[string]Fetcher{
	"keyword": request.TweetsByKeyword,
	"user":    request.TweetsByUser,
}

var searchHandlerMap = map[string]Fetcher{
	"keyword": cache.TweetsByKeyword,
	"user":    cache.TweetsByUser,
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
		end, start := endTime, startTime
		maxTime := time.Now().Add((time.Hour * (24*time.Duration(-nOfDaysAllowed) + 1)) + (time.Minute * time.Duration(59)))
		// fixing the startTime is only worth if the user is acutally interested in
		// the last N_OF_DAYS_ALLOWED days of activity
		if time.Since(startTime).Hours()/24 > nOfDaysAllowed && endTime.After(maxTime) {
			startTime = maxTime
		}
		if time.Now().Sub(endTime).Seconds() < 10 {
			endTime = time.Now().Add(time.Second * time.Duration(-10))
		}

		log.Printf("Querying: \"%s\" %d %v %v", query, nOfAPITweets, startTime, endTime)
		tweets, err = handler1(query, nOfAPITweets, &startTime, &endTime)
		if err != nil {
			log.Printf("Error while querying Twitter: %v", err)
		}
		log.Printf("> Found %d tweets from the API", len(tweets))
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

		tweets, err = handler2(query, uint(amount), &start, &end)
		if err != nil {
			sendError(w, http.StatusInternalServerError, APIError{
				Message: "Could not fetch tweets from cache",
				Error:   err,
			})
			return
		}
		log.Printf("> Found %d tweets from the cache", len(tweets))
		cached = uint(len(tweets) - int(cached))
	}
	sendJSON(w, http.StatusOK, SearchResponse{
		Tweets: tweets,
		Cached: cached,
	})
}
