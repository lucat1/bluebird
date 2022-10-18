package request

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type transportWithHeader struct {
}

type tweet struct {
	id   string
	text string
}
type tweetRawFromRespone struct {
	EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
	Id                  string
	Text                string
}
type metaTweet struct {
	NextToken   string `json:"next_token"`
	ResultCount int    `json:"result_count"`
	NewestId    string `json:"newest_id"`
	OldestId    string `json:"oldest_id"`
}
type responseFromAPI struct {
	Data []tweetRawFromRespone
	Meta metaTweet
}

func (t *transportWithHeader) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Bearer "+bearer)
	return http.DefaultTransport.RoundTrip(req)
}

var bearer string
var client *http.Client

func Init() {
	bearer = os.Getenv("TWITTER_BEARER")
	client = &http.Client{Transport: &transportWithHeader{}}
}

func requestToAPI(url string) []tweet {
	resp, err := client.Get(url)
	if err != nil {
		log.Panic("Errore richiesta twitter")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Errore richiesta twitter")
	}

	var jsonMap responseFromAPI
	err = json.Unmarshal([]byte(string(body)), &jsonMap)
	if err != nil {
		log.Panic("Errore richiesta twitter")
	}
	var tweets []tweet
	for i := 0; i < len(jsonMap.Data); i++ {
		tweets = append(tweets, tweet{id: jsonMap.Data[i].Id, text: jsonMap.Data[i].Text})
	}
	return tweets
}

func recentByKeyword(keyword string) []tweet {
	return requestToAPI("https://api.twitter.com/2/tweets/search/recent?query=" + url.QueryEscape(keyword))
}
func tweetByUserId(userId string) []tweet {
	return requestToAPI("https://api.twitter.com/2/users/" + url.QueryEscape(userId) + "/tweets")
}

func Test() {
	recentByKeyword("#eredita")
	tweetByUserId("2244994945")
}
