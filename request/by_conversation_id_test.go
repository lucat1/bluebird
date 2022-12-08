package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestByConversationID(t *testing.T) {
	SetClient(byConvIDClient)
	tweets, err := TweetsByConversationID("1600882745034428416", 1, nil, nil)
	assert.Nil(t, err, "Expected search not to error")
	assert.NotEmpty(t, tweets, "Expected to find one tweet")
	assert.EqualValues(t, tweets[0].ID, "1600883734538244097", "Tweet ID different than expected")
}
