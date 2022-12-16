package request

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUsers(t *testing.T) {
	users := byKeywordResponse.Users()
	if len(users) <= 0 {
		t.Errorf("Expected a non-empty users list")
	}
	for _, u := range byKeywordResponse.Includes.Users {
		user, has := users[u.ID]
		if !has {
			t.Errorf("Expected to find decoded user with id %s", u.ID)
		}
		assert.Equal(t, user.ID, u.ID, "Expected IDs to match")
		assert.Equal(t, user.Name, u.Name, "Expected Names to match")
		assert.Equal(t, user.Username, u.Username, "Expected Usernames to match")
		assert.Equal(t, user.ProfileImage, u.ProfileImageURL, "Expected ProfileImages to match")
	}
}

func TestTweets(t *testing.T) {
	tweets, err := byKeywordResponse.Tweets()
	if err != nil {
		t.Errorf("Did not expect an error during conversion: %v", err)
	}
	if len(tweets) <= 0 {
		t.Errorf("Expected a non-empty users list")
	}
	for i, tw := range byKeywordResponse.Data {
		tweet := tweets[i]
		assert.Equal(t, tweet.ID, tw.ID, "Expected IDs to match")
		assert.Equal(t, tweet.Text, tw.Text, "Expected Texts to match")
		assert.Equal(t, tweet.User.ID, tw.AuthorID, "Expected User.ID to match")
		if tweet.Geo != nil {
			assert.Equal(t, tweet.Geo.ID, tw.Geo.PlaceID, "Expected Geo.PlaceID to match")
		} else {
			assert.Nil(t, tw.Geo, "Expected original Tweet's Geo to be null")
		}
		assert.Equal(t, tweet.CreatedAt, tw.CreatedAt, "Expected User.CreatedAt to match")
	}
}

func TestPlaces(t *testing.T) {
	places := byKeywordGeoResponse.Places()
	assert.Equal(t, len(places), len(byKeywordGeoResponse.Includes.Places), "More places than expected")

	for _, pl := range byKeywordGeoResponse.Includes.Places {
		place, has := places[pl.ID]
		assert.True(t, has, "'has' not true as expected")

		var loc Coordinates
		if pl.Geo.Type == "Point" {
			loc = Coordinates(pl.Geo.Coordinates)
		} else {
			loc = Coordinates(pl.Geo.BoundingBox)
		}
		assert.Equal(t, pl.ID, place.ID, "Place ID not as expected")
		assert.Equal(t, pl.Geo.Type, place.Type, "Geo Type not as expected")
		assert.Equal(t, loc, place.Coordinates, "Place coordinates not as expected")
	}
}

func TestCoordinates(t *testing.T) {
	coord := &Coordinates{}
	expCoordStr := "[0.3,0.4,0.5,0.6]"
	expCoord := Coordinates{0.3, 0.4, 0.5, 0.6}
	coord.Scan(expCoordStr)
	assert.Equal(t, expCoord, *coord, "Expected coordinates should be [0.3,0.4,0.5,0.6]")
	coordStr, _ := coord.Value()
	assert.Equal(t, expCoordStr, coordStr, "Expected coordinates string should be '0.3,0.4,0.5,0.6'")
}

func TestSentiments(t *testing.T) {
	s1 := Sentiment{Label: "s1", Score: 0.1}
	s2 := Sentiment{Label: "s2", Score: 0.2}
	s3 := Sentiment{Label: "s3", Score: 0.3}
	s4 := Sentiment{Label: "s4", Score: 0.4}
	sentiment := &Sentiments{}
	expSentimentStr := `[{"label":"s1","score":0.1},{"label":"s2","score":0.2},{"label":"s3","score":0.3},{"label":"s4","score":0.4}]`
	expSentiment := Sentiments{s1, s2, s3, s4}
	sentiment.Scan(expSentimentStr)
	assert.Equal(t, expSentiment, *sentiment, "Expected sentiments should be s1:0.1, s2:0.2, s3:0.3, s4:0.4")
	sentimentStr, _ := sentiment.Value()
	assert.Equal(t, expSentimentStr, sentimentStr, `Expected sentiments string should be '[{"label":"s1","score":0.1},{"label":"s2","score":0.2},{"label":"s3","score":0.3},{"label":"s4","score":0.4}]'`)
}

func TestMedias(t *testing.T) {
	s1 := Media{MediaKey: "m1", URL: "url1"}
	medias := &Medias{}
	expMediasStr := `[{"media_key":"m1","url":"url1"}]`
	expMedias := Medias{s1}
	medias.Scan(expMediasStr)
	assert.Equal(t, expMedias, *medias, "Expected medias should be [MediaKey:'m1',URL:'url1']")
	mediasStr, _ := medias.Value()
	assert.Equal(t, expMediasStr, mediasStr, `Expected medias string should be '[{"media_key":"m1","url":"url1"}]'`)
}
