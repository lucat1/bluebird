package cache

import (
	"encoding/json"
	"fmt"

	"git.hjkl.gq/bluebird/bluebird/request"
	"github.com/ostafen/clover/v2"
)

const colTweets = "col-tweets"

type TweetField string

const (
	TweetFieldText TweetField = "Text"
	TweetFieldID              = "ID"
)

var DB clover.DB

func InitCache() {
	tmpDB, err := createDB()
	DB = *tmpDB
	if err != nil {
		panic(err)
	}
	has, err := DB.HasCollection(colTweets)
	if err != nil {
		panic(err)
	}
	if !has {
		err = DB.CreateCollection(colTweets)
	}
	all, _ := FindTweetByUsername("salvinimi")
	fmt.Println(len(all))
}

func createDB() (*clover.DB, error) {
	return clover.Open("swe-DB")
}

func docToStruct[T any](doc *clover.Document, el *T) error {
	docMap := (doc.ToMap())
	ss, _ := json.Marshal(docMap)
	err := json.Unmarshal(ss, &el)
	if err != nil {
		return err
	}
	return nil
}
func docArrayToStruct[T any](docs []*clover.Document, arr []T) error {
	for i, doc := range docs {
		err := docToStruct(doc, &arr[i])
		if err != nil {
			return err
		}
	}
	return nil
}

func AddTweet(tweet request.Tweet) (string, error) {
	doc := clover.NewDocumentOf(tweet)
	return DB.InsertOne(colTweets, doc)
}
func AddTweets(tweets []request.Tweet) ([]string, error) {
	IDs := []string{}
	for _, t := range tweets {
		doc := clover.NewDocumentOf(t)
		ID, err := DB.InsertOne(colTweets, doc)
		if err != nil {
			return nil, err
		}
		IDs = append(IDs, ID)
	}
	return IDs, nil
}

func findTweetBy(field string, filter string) ([]request.Tweet, error) {
	docs, err := DB.FindAll(clover.NewQuery(colTweets).Where(clover.Field(string(field)).Like(filter)))
	if err != nil {
		return nil, nil
	}
	arr := make([]request.Tweet, len(docs))
	docArrayToStruct(docs, arr)

	return arr, nil
}

func FindTweetByText(filter string) ([]request.Tweet, error) {
	return findTweetBy(string(TweetFieldText), filter)
}
func FindTweetByID(filter string) ([]request.Tweet, error) {
	return findTweetBy(string(TweetFieldID), filter)
}
func FindTweetByUsername(filter string) ([]request.Tweet, error) {
	docs, err := DB.FindAll(clover.NewQuery(colTweets))
	if err != nil {
		return nil, nil
	}
	arr := make([]request.Tweet, len(docs))
	docArrayToStruct(docs, arr)
	res := []request.Tweet{}
	for _, el := range arr {
		if el.User.Username == filter {
			res = append(res, el)
		}
	}

	return res, nil
}
