package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

var (
	sentimentURL *url.URL
)

func SetSentimentURL(u string) (err error) {
	if u == "" {
		return nil
	}
	sentimentURL, err = url.Parse(u)
	if err != nil {
		return
	}
	return
}

func SentimentsFromTweets(tweets []Tweet) (sentiments map[string]Sentiments, err error) {
	if sentimentURL == nil {
		return
	}
	u := &url.URL{
		Path: "/predict/multi",
	}

	tweetsJSON, err := json.Marshal(tweets)
	if err != nil {
		return
	}
	res, err := http.Post(sentimentURL.ResolveReference(u).String(), "application/json", bytes.NewBuffer(tweetsJSON))
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		return sentiments, fmt.Errorf("Non 200 status code (was %d): %s", res.StatusCode, string(body))
	}
	return sentiments, json.Unmarshal(body, &sentiments)
}

func SentimentsFromTweet(tweet Tweet) (sentiments Sentiments, err error) {
	if sentimentURL == nil {
		return
	}
	u := &url.URL{
		Path: "/predict",
	}
	query := u.Query()
	queryAdd(query, "text", tweet.Text)
	u.RawQuery = query.Encode()

	res, err := http.Get(sentimentURL.ResolveReference(u).String())
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		return sentiments, fmt.Errorf("Non 200 status code (was %d): %s", res.StatusCode, string(body))
	}
	return sentiments, json.Unmarshal(body, &sentiments)
}
