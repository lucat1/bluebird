package cache

import (
	"os"
	"strings"
	"testing"

	"git.hjkl.gq/bluebird/bluebird/request"
	"git.hjkl.gq/bluebird/bluebird/test"
	"github.com/stretchr/testify/assert"
)

var testTweets []request.Tweet

func TestMain(m *testing.M) {
	if err := Open(":memory:"); err != nil {
		panic(err)
	}
	test.ReadJSON("../mock/db_tweets.json", &testTweets)
	for _, tweet := range testTweets {
		tweet.CreatedAt = tweet.CreatedAt.Local()
	}
	code := m.Run()
	if err := Close(); err != nil {
		panic(err)
	}
	os.Exit(code)
}

func TestInsertTweets(t *testing.T) {
	assert.Nil(t, InsertTweets([]request.Tweet{}), "Expected InsertTweets not to error with an empty input")
	assert.Nil(t, InsertTweets(testTweets[:1]), "Expected InsertTweets not to error with a single input")
	tweet, err := TweetByID(testTweets[0].ID)
	assert.Nil(t, err, "Expected to find the newly inserted tweet")
	assert.EqualValues(t, tweet, testTweets[0], "Unexpected different tweets")
	assert.Nil(t, InsertTweets(testTweets[1:]), "Expected InsertTweets not to error with a lenghty input")
	for i := 1; i < len(testTweets); i++ {
		tweet, err := TweetByID(testTweets[i].ID)
		assert.Nil(t, err, "Expected to find the newly inserted tweet (num %d)", i)
		assert.EqualValues(t, tweet, testTweets[i], "Unexpected different tweets (num %d)", i)
	}
	assert.Nil(t, db.Migrator().DropTable(&request.Tweet{}, &request.User{}, &request.Geo{}), "Failed to clean the Database")
}
func TestTweetsAll(t *testing.T) {
	assert.Nil(t, InsertTweets(testTweets), "Expected InsertTweets not to error with an empty input")
	tweets, err := TweetsAll()
	assert.Nil(t, err, "Failed to load the whole amount of tweets")
	assert.EqualValues(t, tweets, testTweets, "The whole amount of tweets loaded is not the same")
	assert.Nil(t, db.Migrator().DropTable(&request.Tweet{}, &request.User{}, &request.Geo{}), "Failed to clean the Database")
}

func TestTweetsCount(t *testing.T) {
	assert.Nil(t, InsertTweets(testTweets), "Expected InsertTweets not to error with an empty input")
	numOfTweets, err := TweetsCount()
	assert.Nil(t, err, "Failed to load the number of tweets")
	assert.EqualValues(t, len(testTweets), numOfTweets, "The number of tweets is not the same")
	assert.Nil(t, db.Migrator().DropTable(&request.Tweet{}, &request.User{}, &request.Geo{}), "Failed to clean the Database")
}

func TestTweetsByKeyword(t *testing.T) {
	const testString string = "ciao"
	assert.Nil(t, InsertTweets(testTweets), "Expected InsertTweets not to error with an empty input")
	tweets, err := TweetsByKeyword(testString, 20)
	assert.Nil(t, err, "Failed to load tweets using TweetsByKeyword")
	for i, tweet := range tweets {
		assert.True(t, strings.Contains(tweet.Text, testString), "Tweet number %d doesn't contain the searched keyword", i)
	}
	assert.Nil(t, db.Migrator().DropTable(&request.Tweet{}, &request.User{}, &request.Geo{}), "Failed to clean the Database")
}
