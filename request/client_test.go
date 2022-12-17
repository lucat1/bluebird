package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestV1Client(t *testing.T) {
	consumerKey := "fake"
	consumerSecret := "fake"
	oauthToken := "fake"
	oauthSecret := "fake"
	c, err := NewV1Client("https://upload.twitter.com/1.1/media/upload.json?media_category=tweet_image", "https://api.twitter.com/2/", consumerKey, consumerSecret, oauthToken, oauthSecret)
	assert.Nil(t, err, "Expected no error")
	SetV1Client(c)
	assert.Equal(t, c, v1Client, "SetV1Client should set a new client")
	assert.Equal(t, client, Client(), "Client() should return the actual client")
	assert.Equal(t, c, V1Client(), "V1Client() should return the actual v1client")
}
