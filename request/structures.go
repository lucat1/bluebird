package request

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Tweet struct {
	ID        string    `json:"id" gorm:"primaryKey;uniqueIndex"`
	Text      string    `json:"text"`
	UserID    string    `json:"-"`
	User      User      `json:"user"`
	GeoID     *string   `json:"-"`
	Geo       *Geo      `json:"geo"`
	CreatedAt time.Time `json:"created_at" sql:"type:timestamp with time zone"`
}

type Geo struct {
	ID          string      `json:"id" gorm:"primaryKey;uniqueIndex"`
	Coordinates Coordinates `json:"coordinates" gorm:"type:text"`
	PlaceID     string      `json:"place_id" gorm:"primaryKey;uniqueIndex"`
}
type Coordinates []float64

func (sla *Coordinates) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), sla)
}

func (sla Coordinates) Value() (driver.Value, error) {
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
	CreatedAt       string         `json:"created_at"`
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
	Geo                 *rawGeo
}

type rawGeo struct {
	Coordinates rawCoordinates
	PlaceID     string `json:"place_id"`
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
	Users []rawUser
}

type tweetResponse struct {
	Data     []rawTweet
	Meta     metaTweet
	Includes includesTweet
}

func (res *tweetResponse) Users() map[string]User {
	users := map[string]User{}
	for _, u := range res.Includes.Users {
		user := User{ID: u.ID, Name: u.Name, Username: u.Username, ProfileImage: u.ProfileImageURL}
		users[u.ID] = user
	}
	return users
}

func (res *tweetResponse) Tweets() ([]Tweet, error) {
	tweets := []Tweet{}
	users := res.Users()
	for _, t := range res.Data {
		if _, has := users[t.AuthorID]; !has {
			return tweets, fmt.Errorf("User with id %s is not included in Twitter's response", t.AuthorID)
		}

		var (
			geoID *string = nil
			geo   *Geo    = nil
		)
		if t.Geo != nil && t.Geo.Coordinates.Type == "Point" {
			geoID = &t.Geo.PlaceID
			geo = &Geo{
				Coordinates: t.Geo.Coordinates.Coordinates,
				PlaceID:     t.Geo.PlaceID,
			}

		}

		tweets = append(tweets, Tweet{
			ID:        t.ID,
			Text:      t.Text,
			UserID:    t.AuthorID,
			User:      users[t.AuthorID],
			CreatedAt: t.CreatedAt,
			GeoID:     geoID,
			Geo:       geo,
		})
	}
	return tweets, nil
}

func (res *userResponse) User() User {
	return User{ID: res.Data.ID, Name: res.Data.Name, Username: res.Data.Username, ProfileImage: res.Data.ProfileImageURL}
}
