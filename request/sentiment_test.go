package request

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRandomSentiment (t *testing.T) {
  A := randomSentiment()
  assert.Equal(t,A[0]+A[1]+A[2]+A[3],float32(1),"Random Sentiment Values do not add up to 1")
}

func TestSentimentsFromTweet(t *testing.T){
	tweet := Tweet{
		Text: "Ciao mamma sto testando :D",
	}
	retValue, err := SentimentsFromTweet(tweet)
	if err != nil{
	assert.EqualValues(t, rawSentimentResponse.Sentiments, retValue, "Returned sentiment is not as expected")
	}
}