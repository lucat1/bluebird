package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomSentiment(t *testing.T) {
	A := randomSentiment()
	assert.Equal(t, A[0]+A[1]+A[2]+A[3], float32(1), "Random Sentiment Values do not add up to 1")
}

func TestSentimentsFromTweet(t *testing.T) {
	tweet := Tweet{
		Text: "Ciao mamma sto testando :D",
	}
	retValue, err := SentimentsFromTweet(tweet)
	if err != nil {
		assert.EqualValues(t, rawSentimentResponse.Sentiments, retValue, "Returned sentiment is not as expected")
	}
	sentimentURL = nil
	_, err = SentimentsFromTweet(tweet)
	assert.Nil(t, err, "No error expected, sentiments should be generated randomly")
}

func TestSetSentimentURL(t *testing.T) {
	err := SetSentimentURL("")
	assert.Nil(t, err, "No error expected")
	err = SetSentimentURL(string([]byte{00}))
	assert.NotNil(t, err, "Error expected")
}
