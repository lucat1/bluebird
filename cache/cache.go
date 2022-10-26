package cache

import (
	"git.hjkl.gq/bluebird/bluebird/request"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TweetField string

const (
	TweetFieldText      TweetField = "text"
	TweetFieldID                   = "id"
	TweetFieldCreatedAt            = "created_at"
)

var db *gorm.DB

func Open() (err error) {
	if db, err = gorm.Open(sqlite.Open("bluebird.db"), &gorm.Config{}); err != nil {
		return
	}
	db.AutoMigrate(&request.Tweet{}, &request.User{}, &request.Geo{})

	return
}

func Close() error {
	closingDB, err := db.DB()
	if err != nil {
		return err
	}
	return closingDB.Close()
}

func InsertTweets(tweets []request.Tweet) error {
	return db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&tweets).Error
}

func TweetsAny() (res []request.Tweet, _ error) {
	return res, db.Find(&res).Error
}

func TweetsByKeyword(filter string) (res []request.Tweet, _ error) {
	return res, db.Where("text LIKE ?", "%"+filter+"%").Find(&res).Error
}

func TweetByID(filter string) (res request.Tweet, _ error) {
	return res, db.First(&res, request.Tweet{ID: filter}).Error
}

func TweetsByUser(filter string) (res []request.Tweet, err error) {
	err = db.Find(&res, request.Tweet{User: request.User{Username: filter}}).Error
	filtered := []request.Tweet{}
	for _, el := range res {
		if el.User.Username != filter {
			filtered = append(filtered, el)
		}
	}

	return filtered, nil
}
