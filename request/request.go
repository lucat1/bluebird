package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type tweet struct {
	ID   string
	Text string
	User user
	Geo  string
}
type user struct {
	ID           string
	Name         string
	Username     string
	ProfileImage string
}

type userRawMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}
type userRawEntitiesURL struct {
	Start       int
	End         int
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
}
type userRawEntitiesURLs struct {
	URLs userRawEntitiesURL `json:"urls"`
}
type userRawEntities struct {
	URL userRawEntitiesURLs `json:"url"`
}
type userRawResponse struct {
	ID              string `json:"id"`
	Name            string
	Username        string
	URL             string
	Description     string
	ProfileImageURL string `json:"profile_image_url"`
	Verified        bool
	Protected       bool
	CreatedAt       string         `json:"created_at"`
	PublicMetrics   userRawMetrics `json:"public_metrics"`
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
type includesTweet struct {
	Users []userRawResponse
}
type responseFromTweetAPI struct {
	Data     []tweetRawResponse
	Meta     metaTweet
	Includes includesTweet
}

func requestToTweetAPI(link string, count int) ([]tweet, error) {
	var tweets []tweet
	parsedUrl, err := url.Parse(link)
	queries := parsedUrl.Query()
	queries.Add("tweet.fields", "author_id,geo")
	queries.Add("user.fields", "created_at,description,entities,id,location,name,pinned_tweet_id,profile_image_url,protected,public_metrics,url,username,verified,withheld")
	queries.Add("expansions", "author_id")
	parsedUrl.RawQuery = queries.Encode()
	link = parsedUrl.String()
	fmt.Println(link)
	if err != nil {
		return tweets, err
	}
	for len(tweets) < count {
		resp, err := client.Get(link)
		if err != nil {
			return tweets, err
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return tweets, err
		}

		var jsonMap responseFromTweetAPI
		err = json.Unmarshal([]byte(string(body)), &jsonMap)
		if err != nil {
			return tweets, err
		}
		usersMap := map[string]user{}
		for _, u := range jsonMap.Includes.Users {
			user := user{ID: u.ID, Name: u.Name, Username: u.Username, ProfileImage: u.ProfileImageURL}
			usersMap[u.ID] = user
		}
		for _, t := range jsonMap.Data {
			tweets = append(tweets,
				tweet{ID: t.ID, Text: t.Text, User: usersMap[t.AuthorID], Geo: t.Geo})
		}
		if jsonMap.Meta.ResultCount == 0 || jsonMap.Meta.NextToken == "" {
			break
		}
		queries := parsedUrl.Query()
		queries.Set("pagination_token", jsonMap.Meta.NextToken)
		parsedUrl.RawQuery = queries.Encode()
		link = parsedUrl.String()
	}
	return tweets, nil
}

func requestToUserAPI(link string) (user, error) {
	var userResp user
	resp, err := client.Get(link)
	if err != nil {
		return userResp, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return userResp, err
	}

	var jsonMap responseFromUserAPI
	err = json.Unmarshal([]byte(string(body)), &jsonMap)
	if err != nil {
		return userResp, err
	}
	userResp = user{ID: jsonMap.Data.ID, Name: jsonMap.Data.Name, Username: jsonMap.Data.Username}
	return userResp, nil
}

func recentsByKeyword(keyword string, count int) ([]tweet, error) {
	return requestToTweetAPI("https://api.twitter.com/2/tweets/search/recent?query="+url.QueryEscape(keyword), count)
}
func tweetsByUserID(userID string, count int) ([]tweet, error) {
	return requestToTweetAPI("https://api.twitter.com/2/users/"+url.QueryEscape(userID)+"/tweets", count)
}
func tweetsByUsername(username string, count int) ([]tweet, error) {
	userID, err := userIDByUsername(username)
	if err != nil {
		return nil, err
	}
	return tweetsByUserID(userID, count)
}
func userIDByUsername(username string) (string, error) {
	user, err := requestToUserAPI("https://api.twitter.com/2/users/by/username/" + url.QueryEscape(username))
	return user.ID, err
}

func Test() {
	fmt.Println(recentsByKeyword("#eredita", 24))
	fmt.Println(tweetsByUserID("2244994945", 10))
	fmt.Println(tweetsByUsername("matteosalvinimi", 20))
}
