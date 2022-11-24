package server

import (
	"net/http"

	"git.hjkl.gq/team14/team14/cache"
	"git.hjkl.gq/team14/team14/request"
)

type SentimentResponse struct {
	Sentiments request.Sentiments `json:"sentiments"`
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
	sendJSON(w, http.StatusOK, SentimentResponse{
		Sentiments: sentiments,
	})
}
