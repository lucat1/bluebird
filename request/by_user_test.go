package request

import (
	// "github.com/stretchr/testify/assert"
	"testing"
)

func TestByUser(t *testing.T) {
	const l = 5
	SetClient(byUserClient)
	tweets, err := TweetsByUser("", l)
	if err != nil {
		t.Errorf("Expected no error, got %e", err)
	}
	if len(tweets) != 5 {
		t.Errorf("Expected the amount of tweets to be as requested")
	}
	// tws, err := byUserResponse.Tweets()
	// if err != nil {
	// 	t.Errorf("Did not expect error while decoding sample tweet data")
	// }
	// assert.Equal(t, tweets, tws[:l], "Expected IDs to match")
}
