package request

import (
	"fmt"
	"time"
)

func Replies(tweetID string, n uint, startTime, endTime *time.Time) (tweets []Tweet, err error) {
	t, err := TweetByTweetID(tweetID, 1, nil, nil)
	if err != nil {
		return nil, err
	}
	if len(t) < 1 {
		return nil, fmt.Errorf("Tweet not found")
	}
	tweets, err = TweetsByConversationID(t[0].ConversationID, n, startTime, endTime)
	if err != nil {
		return nil, err
	}
	return
}
