package cache

import (
	"log"
	"os"
	"time"

	"git.hjkl.gq/bluebird/bluebird/request"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

type TweetField string

const (
	TweetFieldText      TweetField = "text"
	TweetFieldID                   = "id"
	TweetFieldCreatedAt            = "created_at"
)

var db *gorm.DB

func Open(path string, level logger.LogLevel) (err error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  level,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	if db, err = gorm.Open(sqlite.Open(path), &gorm.Config{
		Logger: newLogger,
	}); err != nil {
		return
	}
	return db.AutoMigrate(&request.Tweet{}, &request.User{}, &request.Geo{})
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

func TweetsByKeyword(filter string, n uint, startTime string, endTime string) (res []request.Tweet, err error) {
	query := db.Limit(int(n)).Preload("User").Preload("Geo")
	if startTime != "" && endTime != "" {
		err = query.Find(&res, "text LIKE ? AND created_at >= ? AND created_at <= ?", filter, startTime, endTime).Error
	} else {
		err = query.Find(&res, "text LIKE ?", filter).Error
	}
	return res, err
}

func TweetByID(filter string) (res request.Tweet, _ error) {
	return res, db.Preload("User").Preload("Geo").First(&res, request.Tweet{ID: filter}).Error
}

func TweetsByUser(username string, n uint, startTime string, endTime string) (res []request.Tweet, err error) {
	query := db.Joins("INNER JOIN users ON users.id = tweets.user_id AND users.username = ?", username).Limit(int(n)).Preload("User").Preload("Geo")
	if startTime != "" && endTime != "" {
		err = query.Find(&res, "created_at >= ? AND created_at <= ?", startTime, endTime).Error
	} else {
		err = query.Find(&res).Error
	}
	return
}
