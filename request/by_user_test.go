package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByUser(t *testing.T) {
	const l = 5
	SetClient(byUserClient)
	tweets, err := TweetsByUser("salvinimi", l)
	assert.Nil(t, err, "Expected TweetsByuser not to error")
	assert.Equal(t, len(tweets), l, "Expected the amount of tweets to be as required")
	twts, err := byUserResponse.Tweets()
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, tweets, twts[:l], "Expected IDs to match")
}
