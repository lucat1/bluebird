package cache

import (
	"git.hjkl.gq/bluebird/bluebird/request"
)

type TweetField string

const (
	TweetFieldText      TweetField = "text"
	TweetFieldID                   = "id"
	TweetFieldCreatedAt            = "created_at"
)

var db *bbolt.DB

func Open() (err error) {
	if db, err = clover.Open("bluebird.db"); err != nil {
		return
	}
	return initialize()
}

func initialize() (err error) {
	/// init tables n stuff
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

func InsertTweets(tweets []request.Tweet) (IDs []string, err error) {
	tx, err := db.Begin(true)
	if err != nil {
		return
	}
	defer tx.Rollback()
	tweetsById := tx.Bucket(tweetsByID)

	for tweet := range tweets {
		_, err = tx.CreateBucket(bucket)
		if err != nil {
			return
		}
	}

	return tx.Commit()
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

func TweetsAny() (res []request.Tweet, err error) {
	docs, err := db.FindAll(clover.NewQuery(tweetsCollection))
	if err != nil {
		return nil, nil
	}
	if res, err = multiUnmarshal[request.Tweet](docs); err != nil {
		return
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
