package cache

import (
	"os"
	"strings"
	"testing"

	"git.hjkl.gq/bluebird/bluebird/request"
	"git.hjkl.gq/bluebird/bluebird/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm/logger"
)

var testTweets []request.Tweet

func clearDB() (err error) {
	err = db.Where("1 = 1").Delete(&request.Tweet{}).Error
	if err != nil {
		return
	}
	err = db.Where("1 = 1").Delete(&request.User{}).Error
	if err != nil {
		return
	}
	err = db.Where("1 = 1").Delete(&request.Geo{}).Error
	if err != nil {
		return
	}
	return
}

func TestMain(m *testing.M) {
	if err := Open(":memory:", logger.Silent); err != nil {
		panic(err)
	}
	test.ReadJSON("../mock/db_tweets.json", &testTweets)
	for i := range testTweets {
		testTweets[i].CreatedAt = testTweets[i].CreatedAt.UTC()
	}
	code := m.Run()
	if err := Close(); err != nil {
		panic(err)
	}
	os.Exit(code)
}

func TestInsertTweets(t *testing.T) {
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
	assert.Nil(t, clearDB(), "Failed to clean the Database")
}

func TestTweetsAll(t *testing.T) {
	assert.Nil(t, InsertTweets(testTweets), "Expected InsertTweets not to error while filling in test data")
	tweets, err := TweetsAll()
	assert.Nil(t, err, "Failed to load the whole amount of tweets")
	assert.Equal(t, len(tweets), len(testTweets), "The amount ot tweets matches the inserted payload")
	assert.EqualValues(t, tweets, testTweets, "The whole amount of tweets loaded is not the same")
	assert.Nil(t, clearDB(), "Failed to clean the Database")
}

func TestTweetsCount(t *testing.T) {
	assert.Nil(t, InsertTweets(testTweets), "Expected InsertTweets not to error while filling in test data")
	numOfTweets, err := TweetsCount()
	assert.Nil(t, err, "Failed to load the number of tweets")
	assert.EqualValues(t, len(testTweets), numOfTweets, "The number of tweets is not the same")
	assert.Nil(t, clearDB(), "Failed to clean the Database")
}

func TestTweetsByKeyword(t *testing.T) {
	const testString string = "ciao"
	assert.Nil(t, InsertTweets(testTweets), "Expected InsertTweets not to error while filling in test data")
	tweets, err := TweetsByKeyword(testString, 20, "", "")
	assert.Nil(t, err, "Failed to load tweets using TweetsByKeyword")
	for i, tweet := range tweets {
		assert.True(t, strings.Contains(strings.ToLower(tweet.Text), testString), "Tweet number %d doesn't contain the searched keyword", i)
	}
	assert.Nil(t, clearDB(), "Failed to clean the Database")
}

func TestTweetByID(t *testing.T) {
	const testID string = "1585996206379077635"
	assert.Nil(t, InsertTweets(testTweets), "Expected InsertTweets not to error while filling in test data")
	_, err := TweetByID("invalid_id")
	assert.NotNil(t, err, "Expected TweetById to error with an invalid ID")
	tweet, err := TweetByID(testID)
	assert.Nil(t, err, "Expected TweetById not to error with a valid ID")
	assert.EqualValues(t, tweet, testTweets[0], "Expected the tweet queried by ID to match the inserted data")
	tweet1, err := TweetByID(testID)
	assert.Nil(t, err, "Expected TweetById not to error with a valid ID (2)")
	assert.EqualValues(t, tweet, tweet1, "Expected the tweets queried by ID multiple times to match")
	assert.Nil(t, clearDB(), "Failed to clean the Database")
}

func TestTweetsByUser(t *testing.T) {
	const testUser string = "_ultimotiamo_"
	assert.Nil(t, InsertTweets(testTweets), "Expected InsertTweets not to error while filling in test data")
	tweets, err := TweetsByUser("invalid_user", 50, "", "")
	assert.Nil(t, err, "Expected TweetsByUser not to error with an invalid username")
	assert.Equal(t, len(tweets), 0, "Expected TweetsByUser to return an empty slice with an invalid input")
	tweets, err = TweetsByUser(testUser, 50, "", "")
	assert.Nil(t, err, "Expected TweetsByUser not to error with a valid user")
	assert.Equal(t, len(tweets), 1, "Expected to have found only one tweet")
	assert.EqualValues(t, tweets[0], testTweets[0], "Expected the tweet retrieved by username to match the source one")
	assert.Nil(t, clearDB(), "Failed to clean the Database")
}
