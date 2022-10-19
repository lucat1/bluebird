package request

import (
	"os"
	"testing"

	"git.hjkl.gq/bluebird/bluebird/test"
	"github.com/stretchr/testify/assert"
)

var response tweetResponse

func TestMain(m *testing.M) {
	test.ReadJSON("../mock/by_keyword.json", &response)
	code := m.Run()
	os.Exit(code)
}

func TestUsers(t *testing.T) {
	users := response.Users()
	if len(users) <= 0 {
		t.Errorf("Expected a non-empty users list")
	}
	for _, u := range response.Includes.Users {
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
