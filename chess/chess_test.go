package chess

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMoveFromText(t *testing.T) {
	invalid := "d3"
	assert.Nil(t, moveFromText(invalid), "Expected invalid tweet to return nil")
	valid := "@bluebird_swe d3"
	assert.NotNil(t, moveFromText(valid), "Expected invalid tweet to return nil")
	mv := "d3"
	assert.EqualValues(t, &mv, moveFromText(valid), "Expected valid tweet to return d3")
}

func TestFetchTweets(t *testing.T) {
	m := NewMatch(time.Minute)
	m.move("d3")
	<-time.After(time.Second * (ReplyPollInterval + 1))
	assert.Len(t, m.Tweets, 1, "Expected one tweet to have been fetched")
	m.end()
}
