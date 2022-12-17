package cache

import (
	"os"
	"strings"
	"testing"
	"time"

	"git.hjkl.gq/team14/team14/request"
	"git.hjkl.gq/team14/team14/test"
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
	for i := range testTweets {
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
	tweets, err := TweetsByKeyword(testString, 20, nil, nil)
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
	tweets, err := TweetsByUser("invalid_user", 50, nil, nil)
	assert.Nil(t, err, "Expected TweetsByUser not to error with an invalid username")
	assert.Equal(t, len(tweets), 0, "Expected TweetsByUser to return an empty slice with an invalid input")
	tweets, err = TweetsByUser(testUser, 50, nil, nil)
	assert.Nil(t, err, "Expected TweetsByUser not to error with a valid user")
	assert.Equal(t, 1, len(tweets), "Expected to have found only one tweet")
	assert.EqualValues(t, tweets[0], testTweets[0], "Expected the tweet retrieved by username to match the source one")
	assert.Nil(t, clearDB(), "Failed to clean the Database")
}

func TestTimeRange(t *testing.T) {
	const testUser string = "_ultimotiamo_"
	start, err := time.Parse(time.RFC3339, "2022-10-24T16:06:04.325830601+02:00")
	assert.Nil(t, err, "Could not parse time")
	end := start.Add(time.Hour * time.Duration(24))
	assert.Nil(t, InsertTweets(testTweets), "Expected InsertTweets not to error while filling in test data")
	tweets, err := TweetsByUser(testUser, 50, &start, &end)
	assert.Nil(t, err, "Expected TweetsByUser not to error with an invalid time range")
	assert.Equal(t, len(tweets), 0, "Expected TweetsByUser to return an empty slice with an invalid time range")

	start, err = time.Parse(time.RFC3339, "2022-10-27T16:06:04.325830601+02:00")
	assert.Nil(t, err, "Could not parse time")
	end, err = time.Parse(time.RFC3339, "2022-10-31T16:06:04.325830601+02:00")
	assert.Nil(t, err, "Could not parse time")
	tweets, err = TweetsByUser(testUser, 50, &start, &end)
	assert.Nil(t, err, "Expected TweetsByUser not to error with a valid time range")
	assert.Equal(t, 1, len(tweets), "Expected to have found only one tweet")
	assert.EqualValues(t, tweets[0], testTweets[0], "Expected the tweet retrieved by username and time range to match the source one")
	assert.Nil(t, clearDB(), "Failed to clean the Database")
}

func TestPoliticians(t *testing.T) {
	p1 := request.Politician{Name: "N1", Surname: "S1", NPosts: 1, Points: 1, Average: 1, BestSingleScore: 1, LastUpdated: time.Now()}
	p2 := request.Politician{Name: "N2", Surname: "S2", NPosts: 2, Points: 2, Average: 2, BestSingleScore: 2, LastUpdated: time.Now()}
	p3 := request.Politician{Name: "N3", Surname: "S3", NPosts: 15, Points: 20, Average: 1.33, BestSingleScore: 4, LastUpdated: time.Now()}
	ps := []request.Politician{p1, p2, p3}
	err := InsertPoliticians(ps)
	assert.Nil(t, err, "Failed to insert politicians")
	res, err := PoliticianByNameSurname(p1.Name, p1.Surname)
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, p1.Name, res.Name, "Same name expected")
	assert.Equal(t, p1.Surname, res.Surname, "Same surname expected")

	scoreboard, err := PoliticiansScoreboard()
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, p3.Name, scoreboard[0].Name, "P3 expected to be first place")
	assert.Equal(t, p3.Surname, scoreboard[0].Surname, "P3 expected to be first place")
	assert.Equal(t, p2.Name, scoreboard[1].Name, "P2 expected to be second place")
	assert.Equal(t, p2.Surname, scoreboard[1].Surname, "P2 expected to be second place")

	res, err = PoliticianBestAverage()
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, p2.Name, res.Name, "P2 expected to had the best average")
	assert.Equal(t, p2.Surname, res.Surname, "P2 expected to had the best average")

	res, err = PoliticianBestSingleScore()
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, p3.Name, res.Name, "P3 expected to had the best single score")
	assert.Equal(t, p3.Surname, res.Surname, "P3 expected to had the best single score")

	addNew := request.Politician{Name: "N4", Surname: "S4", NPosts: 1, Points: 4, Average: 1, BestSingleScore: 4, LastUpdated: time.Now().Add(20)}
	err = AddPointsPoliticianByNameSurname(addNew)
	assert.Nil(t, err, "Expected no error")

	addP1 := request.Politician{Name: "N1", Surname: "S1", NPosts: 1, Points: 4, Average: 1, BestSingleScore: 4, LastUpdated: time.Now().Add(20)}
	err = AddPointsPoliticianByNameSurname(addP1)
	assert.Nil(t, err, "Expected no error")
	res, err = PoliticianByNameSurname(p1.Name, p1.Surname)
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, p1.Points+addP1.Points, res.Points, "Expected points to be added")
	backupPoints := res.Points

	addP1 = request.Politician{Name: "N1", Surname: "S1", NPosts: 1, Points: 4, Average: 1, BestSingleScore: 4, LastUpdated: time.Now().AddDate(0, 0, -1)}
	err = AddPointsPoliticianByNameSurname(addP1)
	assert.Nil(t, err, "Expected no error")
	res, err = PoliticianByNameSurname(p1.Name, p1.Surname)
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, backupPoints, res.Points, "Expected points not to be added")

	addP1 = request.Politician{Name: "N1", Surname: "S1", NPosts: 1, Points: 4, Average: 1, BestSingleScore: 4, LastUpdated: time.Now().Add(40)}
	addP2 := request.Politician{Name: "N2", Surname: "S2", NPosts: 1, Points: 4, Average: 1, BestSingleScore: 4, LastUpdated: time.Now().Add(40)}
	adds := []request.Politician{addP1, addP2}
	err = AddPointsPoliticians(adds)
	assert.Nil(t, err, "Expected no error")
	res, err = PoliticianByNameSurname(p1.Name, p1.Surname)
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, backupPoints+addP1.Points, res.Points, "Expected points to be added")
}

func TestTeams(t *testing.T) {
	t1 := request.Team{Username: "t1", PictureURL: "url1"}
	t2 := request.Team{Username: "t2", PictureURL: "url2"}
	ts := []request.Team{t1, t2}
	err := InsertTeams(ts)
	assert.Nil(t, err, "Expected no error")

	res, err := TeamsAll()
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, ts, res, "Teams should be the same inserted before")

	st1, err := SearchTeamByUsername(t1.Username)
	assert.Nil(t, err, "Expected no error")
	assert.Equal(t, t1, st1, "Expected no error")

	_, err = SearchTeamByUsername("not found")
	assert.NotNil(t, err, "Expected error")
}

func TestClose(t *testing.T) {
	err := Close()
	assert.Nil(t, err, "Expected no error")
}
