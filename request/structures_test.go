package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {
	users := byKeywordResponse.Users()
	if len(users) <= 0 {
		t.Errorf("Expected a non-empty users list")
	}
	for _, u := range byKeywordResponse.Includes.Users {
		user, has := users[u.ID]
		if !has {
			t.Errorf("Expected to find decoded user with id %s", u.ID)
		}
		assert.Equal(t, user.ID, u.ID, "Expected IDs to match")
		assert.Equal(t, user.Name, u.Name, "Expected Names to match")
		assert.Equal(t, user.Username, u.Username, "Expected Usernames to match")
		assert.Equal(t, user.ProfileImage, u.ProfileImageURL, "Expected ProfileImages to match")
	}
}

func TestTweets(t *testing.T) {
	tweets, err := byKeywordResponse.Tweets()
	if err != nil {
		t.Errorf("Did not expect an error during conversion: %v", err)
	}
	if len(tweets) <= 0 {
		t.Errorf("Expected a non-empty users list")
	}
	for i, tw := range byKeywordResponse.Data {
		tweet := tweets[i]
		assert.Equal(t, tweet.ID, tw.ID, "Expected IDs to match")
		assert.Equal(t, tweet.Text, tw.Text, "Expected Texts to match")
		assert.Equal(t, tweet.User.ID, tw.AuthorID, "Expected User.ID to match")
		if tweet.Geo != nil {
			assert.Equal(t, tweet.Geo.PlaceID, tw.Geo.PlaceID, "Expected Geo.PlaceID to match")
			assert.Equal(t, tweet.Geo.Coordinates, tw.Geo.Coordinates.Coordinates, "Expected Geo.Coordinates to match")
		} else {
			assert.Nil(t, tw.Geo, "Expected original Tweet's Geo to be null")
		}
		assert.Equal(t, tweet.CreatedAt, tw.CreatedAt, "Expected User.CreatedAt to match")
	}
}
