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
		t.Errorf("Did not expect an error during conversion: %e", err)
	}
	if len(tweets) <= 0 {
		t.Errorf("Expected a non-empty users list")
	}
	for i, tw := range byKeywordResponse.Data {
		tweet := tweets[i]
		assert.Equal(t, tweet.ID, tw.ID, "Expected IDs to match")
		assert.Equal(t, tweet.Text, tw.Text, "Expected Texts to match")
		assert.Equal(t, tweet.User.ID, tw.AuthorID, "Expected User.ID to match")
		assert.Equal(t, tweet.Geo, tw.Geo, "Expected Geos to match")
	}
}
