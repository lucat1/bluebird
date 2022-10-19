package request

import "fmt"

type Tweet struct {
	ID   string
	Text string
	User User
	Geo  string
}

type User struct {
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

func (res *responseFromTweetAPI) Users() map[string]User {
	users := map[string]User{}
	for _, u := range res.Includes.Users {
		user := User{ID: u.ID, Name: u.Name, Username: u.Username, ProfileImage: u.ProfileImageURL}
		users[u.ID] = user
	}
	return users
}

func (res *responseFromTweetAPI) Tweets() ([]Tweet, error) {
	tweets := []Tweet{}
	users := res.Users()
	for _, t := range res.Data {
		if _, has := users[t.AuthorID]; has {
			return tweets, fmt.Errorf("User with id %s is not included in Twitter's response", t.AuthorID)
		}

		tweets = append(tweets, Tweet{
			ID:   t.ID,
			Text: t.Text,
			User: users[t.AuthorID],
			Geo:  t.Geo,
		})
	}
	return tweets, nil
}
