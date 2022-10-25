package cache

import (
	"git.hjkl.gq/bluebird/bluebird/request"
	"github.com/ostafen/clover/v2"
)

const tweetsCollection = "tweets"

type TweetField string

const (
	TweetFieldText      TweetField = "text"
	TweetFieldID                   = "id"
	TweetFieldCreatedAt            = "created_at"
)

var db *clover.DB

func Open() (err error) {
	if db, err = clover.Open("bluebird_db"); err != nil {
		return
	}
	return initialize()
}

func initialize() (err error) {
	var has bool
	if has, err = db.HasCollection(tweetsCollection); err != nil {
		return
	}
	if !has {
		if err = db.CreateCollection(tweetsCollection); err != nil {
			return
		}
	}
	return
}

func Close() error {
	return db.Close()
}

func multiUnmarshal[T any](docs []*clover.Document) (res []T, err error) {
	var obj T
	for _, doc := range docs {
		if err = doc.Unmarshal(&obj); err != nil {
			return
		}
		res = append(res, obj)
	}
	return
}

func InsertTweet(tweet request.Tweet) (string, error) {
	return db.InsertOne(tweetsCollection, clover.NewDocumentOf(tweet))
}

func InsertTweets(tweets []request.Tweet) (IDs []string, err error) {
	var ID string
	for _, doc := range tweets {
		ID, err = InsertTweet(doc)
		if err != nil {
			return
		}
		IDs = append(IDs, ID)
	}
	return
}

func tweetsBy(field TweetField, filter string) (res []request.Tweet, err error) {
	docs, err := db.FindAll(clover.NewQuery(tweetsCollection).Where(clover.Field(string(field)).Like(filter)))
	if err != nil {
		return nil, nil
	}
	if res, err = multiUnmarshal[request.Tweet](docs); err != nil {
		return
	}
	return
}

func TweetsByKeyword(filter string) ([]request.Tweet, error) {
	return tweetsBy(TweetFieldText, filter)
}

func TweetsByID(filter string) ([]request.Tweet, error) {
	return tweetsBy(TweetFieldID, filter)
}

func TweetsByUser(filter string) (res []request.Tweet, err error) {
	docs, err := db.FindAll(clover.NewQuery(tweetsCollection))
	if err != nil {
		return nil, nil
	}
	if res, err = multiUnmarshal[request.Tweet](docs); err != nil {
		return
	}
	filtered := []request.Tweet{}
	for _, el := range res {
		if el.User.Username != filter {
			filtered = append(filtered, el)
		}
	}

	return filtered, nil
}
