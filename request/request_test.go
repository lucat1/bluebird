package request

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestRaw(t *testing.T) {
	myurl, _ := url.Parse("users/by/username/quizzettone")
	res, err := requestRaw[userResponse](requestRawClient, myurl)
	user := res.User()
	assert.Nil(t, err, "Expected no error")
	assert.EqualValues(t, requestRawResponse.User(), user, "Returned user is not as expected")

	_, err = requestRaw[userResponse](requestRawErrorClient, myurl)
	assert.NotNil(t, err, "Expected an error")
}

func TestRequestUser(t *testing.T) {
	SetClient(requestUserErrorClient)
	myurl, _ := url.Parse("users/by/username/quizzettone")
	_, err := requestUser(myurl)
	assert.NotNil(t, err, "Expected an error")
}

func TestRequestTweets(t *testing.T) {
	SetClient(requestTweetsErrorClient)
	myurl, _ := url.Parse("/users/270839361/tweets")
	_, err := requestTweets(myurl, 20)
	assert.NotNil(t, err, "Expected an error")
}
