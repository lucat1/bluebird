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
	ID       string
	Text     string
	AuthorID string
	Geo      string
}
type user struct {
	ID       string
	Name     string
	Username string
}

type userRawResponse struct {
	ID       string `json:"id"`
	Name     string
	Username string
}
type responseFromUserAPI struct {
	Data userRawResponse
}

type tweetRawResponse struct {
	EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
	ID                  string   `json:"id"`
	Text                string
	AuthorID            string `json:"author_id"`
	Geo                 string
}
type metaTweet struct {
	NextToken   string `json:"next_token"`
	ResultCount int    `json:"result_count"`
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
}
type responseFromTweetAPI struct {
	Data []tweetRawResponse
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

func requestToTweetAPI(link string, count int) []tweet {
	var tweets []tweet
	parsedUrl, err := url.Parse(link)
	queries := parsedUrl.Query()
	queries.Add("tweet.fields", "author_id,geo")
	parsedUrl.RawQuery = queries.Encode()
	link = parsedUrl.String()
	fmt.Println(link)
	if err != nil {
		log.Panic("Errore richiesta twitter")
	}
	for len(tweets) < count {
		resp, err := client.Get(link)
		if err != nil {
			log.Panic("Errore richiesta twitter")
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Panic("Errore richiesta twitter")
		}

		var jsonMap responseFromTweetAPI
		err = json.Unmarshal([]byte(string(body)), &jsonMap)
		if err != nil {
			log.Panic("Errore richiesta twitter")
		}
		for i := 0; i < len(jsonMap.Data) && len(tweets) < count; i++ {
			tweets = append(tweets, tweet{ID: jsonMap.Data[i].ID, Text: jsonMap.Data[i].Text, AuthorID: jsonMap.Data[i].AuthorID, Geo: jsonMap.Data[i].Geo})
		}
		if jsonMap.Meta.ResultCount == 0 || jsonMap.Meta.NextToken == "" {
			break
		}
		queries := parsedUrl.Query()
		queries.Set("pagination_token", jsonMap.Meta.NextToken)
		parsedUrl.RawQuery = queries.Encode()
		link = parsedUrl.String()
	}
	fmt.Println(len(tweets))
	return tweets
}

func requestToUserAPI(link string) user {
	var userResp user
	parsedUrl, err := url.Parse(link)
	queries := parsedUrl.Query()
	queries.Add("tweet.fields", "author_id,geo")
	parsedUrl.RawQuery = queries.Encode()
	link = parsedUrl.String()
	if err != nil {
		log.Panic("Errore richiesta twitter")
	}
	resp, err := client.Get(link)
	if err != nil {
		log.Panic("Errore richiesta twitter")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic("Errore richiesta twitter")
	}

	var jsonMap responseFromUserAPI
	err = json.Unmarshal([]byte(string(body)), &jsonMap)
	if err != nil {
		log.Panic("Errore richiesta twitter")
	}
	userResp = user{ID: jsonMap.Data.ID, Name: jsonMap.Data.Name, Username: jsonMap.Data.Username}
	return userResp
}

func recentsByKeyword(keyword string, count int) []tweet {
	return requestToTweetAPI("https://api.twitter.com/2/tweets/search/recent?query="+url.QueryEscape(keyword), count)
}
func tweetsByUserID(userID string, count int) []tweet {
	return requestToTweetAPI("https://api.twitter.com/2/users/"+url.QueryEscape(userID)+"/tweets", count)
}
func tweetsByUsername(username string, count int) []tweet {
	userID := userIDByUsername(username)
	return requestToTweetAPI("https://api.twitter.com/2/users/"+url.QueryEscape(userID)+"/tweets", count)
}
func userIDByUsername(username string) string {
	return requestToUserAPI("https://api.twitter.com/2/users/by/username/" + url.QueryEscape(username)).ID
}

func Test() {
	// recentsByKeyword("#eredita", 24)
	// tweetsByUserID("2244994945", 49)
	// tweetsByUsername("TwitterDev", 20)
}
