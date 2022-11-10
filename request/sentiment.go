package request

import (
	"math/rand"
	"net/http"
	"net/url"
)

var (
	sentimentClient   *RequestClient
	availableSentimes = []Sentiment{SentimentAnger, SentimentSadness, SentimentFear, SentimentJoy}
)

func SetSentimentURL(u string) (err error) {
	if u == "" {
		return nil
	}
	url, err := url.Parse(u)
	if err != nil {
		return
	}
	sentimentClient = &RequestClient{
		HTTP: http.DefaultClient,
		URL:  url,
	}
	return
}

func SentimentFromText(text string) (_ Sentiment, err error) {
	if sentimentClient == nil {
		return availableSentimes[rand.Intn(4)], nil
	}
	u := &url.URL{
		Path: "/predict",
	}
	u.Query().Add("text", text)
	raw, err := requestRaw[sentimentResponse](sentimentClient, u)
	if err != nil {
		return
	}
	return raw.Sentiment, nil
}
