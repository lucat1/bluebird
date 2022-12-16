package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplies(t *testing.T) {
	SetClient(repliesClient)
	tweets, err := Replies("1600882745034428416", 1, nil, nil)
	assert.Nil(t, err, "Expected search not to error")
	assert.NotEmpty(t, tweets, "Expected to find one tweet")
	assert.EqualValues(t, tweets[0].ID, "1600883734538244097", "Tweet ID different than expected")

	SetClient(repliesNoRepliesClient)
	_, err = Replies("1600882745034428416", 1, nil, nil)
	assert.NotNil(t, err, "Expected search to error")

	SetClient(repliesNoConvClient)
	_, err = Replies("1600882745034428416", 1, nil, nil)
	assert.NotNil(t, err, "Expected search to error")
}
