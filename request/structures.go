package request

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Tweet struct {
	ID             string      `json:"id" gorm:"primaryKey;uniqueIndex"`
	Text           string      `json:"text"`
	UserID         string      `json:"-"`
	User           User        `json:"user"`
	GeoID          *string     `json:"-"`
	Geo            *Geo        `json:"geo"`
	ConversationID string      `json:"conversation_id"`
	CreatedAt      time.Time   `json:"created_at" sql:"type:timestamp with time zone"`
	Sentiments     *Sentiments `json:"sentiments"`
}

type Geo struct {
	ID          string      `json:"id" gorm:"primaryKey;uniqueIndex"`
	Type        string      `json:"type"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates []float64
type Sentiments [4]Sentiment

func (sla *Coordinates) Scan(src interface{}) error {
	return json.Unmarshal([]byte(src.(string)), sla)
}

func (sla Coordinates) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}

func (sla *Sentiments) Scan(src interface{}) error {
	err := json.Unmarshal([]byte(src.(string)), sla)
	return err
}

func (sla Sentiments) Value() (driver.Value, error) {
	val, err := json.Marshal(sla)
	return string(val), err
}

type User struct {
	ID           string `json:"id" gorm:"primaryKey;uniqueIndex"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	ProfileImage string `json:"profile_image"`
}

type userRawMetrics struct {
	FollowersCount int `json:"followers_count"`
	FollowingCount int `json:"following_count"`
	TweetCount     int `json:"tweet_count"`
	ListedCount    int `json:"listed_count"`
}

type rawUserEntitiesURL struct {
	Start       int
	End         int
	URL         string `json:"url"`
	ExpandedURL string `json:"expanded_url"`
	DisplayURL  string `json:"display_url"`
}

type rawUser struct {
	ID              string `json:"id"`
	Name            string
	Username        string
	URL             string
	Description     string
	ProfileImageURL string `json:"profile_image_url"`
	Verified        bool
	Protected       bool
	CreatedAt       time.Time      `json:"created_at"`
	PublicMetrics   userRawMetrics `json:"public_metrics"`
}

type userResponse struct {
	Data rawUser `json:"data"`
}

type rawTweet struct {
	EditHistoryTweetIDs []string `json:"edit_history_tweet_ids"`
	ID                  string   `json:"id"`
	Text                string
	AuthorID            string    `json:"author_id"`
	CreatedAt           time.Time `json:"created_at"`
	Geo                 *struct {
		PlaceID string `json:"place_id"`
	}
	ConversationID string `json:"conversation_id"`
}

type rawPlace struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Geo      rawGeo `json:"geo"`
}

type rawGeo struct {
	Type        string    `json:"type"`
	BoundingBox []float64 `json:"bbox"`
	Coordinates []float64 `json:"coordinates"`
}

type rawCoordinates struct {
	Type        string
	Coordinates []float64
}

type metaTweet struct {
	NextToken   string `json:"next_token"`
	ResultCount int    `json:"result_count"`
	NewestID    string `json:"newest_id"`
	OldestID    string `json:"oldest_id"`
}

type includesTweet struct {
	Users  []rawUser
	Places []rawPlace
}

type tweetResponse struct {
	Data     []rawTweet
	Meta     metaTweet
	Includes includesTweet
}

type MediaResponse struct {
	MediaID string `json:"media_id_string"`
}

type TweetRequest struct {
	Text  string            `json:"text"`
	Media TweetRequestMedia `json:"media"`
}

type TweetRequestMedia struct {
	MediaIDs []string `json:"media_ids"`
}

type rawTweetResponse struct {
	Data TweetResponse `json:"data"`
}

type TweetResponse struct {
	ID string `json:"id"`
}

func (res *tweetResponse) Users() map[string]User {
	users := map[string]User{}
	for _, u := range res.Includes.Users {
		user := User{ID: u.ID, Name: u.Name, Username: u.Username, ProfileImage: u.ProfileImageURL}
		users[u.ID] = user
	}
	return users
}

func (res *tweetResponse) Places() map[string]Geo {
	places := map[string]Geo{}
	for _, p := range res.Includes.Places {
		var loc Coordinates
		if p.Geo.Type == "Point" {
			loc = Coordinates(p.Geo.Coordinates)
		} else {
			loc = Coordinates(p.Geo.BoundingBox)
		}
		place := Geo{
			ID:          p.ID,
			Type:        p.Geo.Type,
			Coordinates: loc,
		}
		places[p.ID] = place
	}
	return places
}

func (res *tweetResponse) Tweets() ([]Tweet, error) {
	tweets := []Tweet{}
	users := res.Users()
	places := res.Places()
	for _, t := range res.Data {
		if _, has := users[t.AuthorID]; !has {
			return tweets, fmt.Errorf("User with id %s is not included in Twitter's response", t.AuthorID)
		}

		var (
			geoID *string = nil
			geo   *Geo    = nil
		)
		if t.Geo != nil {
			geoID = &t.Geo.PlaceID
			g := places[*geoID]
			geo = &g
		}

		tweets = append(tweets, Tweet{
			ID:             t.ID,
			Text:           t.Text,
			UserID:         t.AuthorID,
			User:           users[t.AuthorID],
			CreatedAt:      t.CreatedAt,
			GeoID:          geoID,
			Geo:            geo,
			ConversationID: t.ConversationID,
			Sentiments:     nil,
		})
	}
	return tweets, nil
}

func (res *userResponse) User() User {
	return User{ID: res.Data.ID, Name: res.Data.Name, Username: res.Data.Username, ProfileImage: res.Data.ProfileImageURL}
}

type SentimentType string

const (
	SentimentAnger   SentimentType = "anger"
	SentimentSadness               = "sadness"
	SentimentFear                  = "fear"
	SentimentJoy                   = "joy"
)

type Sentiment struct {
	Label SentimentType `json:"label"`
	Score float32       `json:"score"`
}

type sentimentResponse struct {
	Sentiments Sentiments `json:"sentiment"`
}
