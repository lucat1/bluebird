package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByKeyword(t *testing.T) {
	const l = 5
	SetClient(byKeywordClient)
	tweets, err := TweetsByKeyword("", l, nil, nil)
	assert.Nil(t, err, "Expected TweetsByKeyword not to error")
	assert.Equal(t, len(tweets), l, "Expected the amount of tweets to be as required")
	twts, err := byKeywordResponse.Tweets()
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, tweets, twts[:l], "Expected Tweets to match")
}
