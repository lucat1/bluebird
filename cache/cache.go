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

func Open(path string) (err error) {
	if db, err = gorm.Open(sqlite.Open(path), &gorm.Config{}); err != nil {
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

func TweetsAll() (res []request.Tweet, _ error) {
	return res, db.Preload("User").Preload("Geo").Find(&res).Error
}

func TweetsCount() (n int64, _ error) {
	return n, db.Model(&request.Tweet{}).Count(&n).Error
}

func TweetsByKeyword(filter string, n uint) (res []request.Tweet, _ error) {
	return res, db.Where("text LIKE ?", "%"+filter+"%").Limit(int(n)).Preload("User").Preload("Geo").Find(&res).Error
}

func TweetByID(filter string) (res request.Tweet, _ error) {
	return res, db.Preload("User").Preload("Geo").First(&res, request.Tweet{ID: filter}).Error
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
