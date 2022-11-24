package request

import "fmt"

func Replies(tweetID string, n uint, startTime string, endTime string) (tweets []Tweet, err error) {
	t, err := TweetByTweetID(tweetID, 1, "", "")
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
