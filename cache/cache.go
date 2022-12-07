package cache

import (
	"log"
	"os"
	"strings"
	"time"

	"git.hjkl.gq/team14/team14/request"
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
	return db.AutoMigrate(&request.Tweet{}, &request.User{}, &request.Geo{}, &request.Politician{})
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
		DoNothing: true,
	}).Create(&tweets).Error
}

func TweetsAll() (res []request.Tweet, _ error) {
	return res, db.Preload("User").Preload("Geo").Find(&res).Error
}
func PoliticiansAll() (res []request.Politician, _ error) {
	return res, db.Find(&res).Error
}

func TweetsCount() (n int64, _ error) {
	return n, db.Model(&request.Tweet{}).Count(&n).Error
}

func TweetsByKeyword(filter string, n uint, startTime, endTime *time.Time) (res []request.Tweet, err error) {
	query := db.Limit(int(n)).Preload("User").Preload("Geo")
	if startTime != nil && endTime != nil {
		err = query.Find(&res, "text LIKE ? AND created_at >= ? AND created_at <= ?", "%"+filter+"%", startTime, endTime).Error
	} else {
		err = query.Find(&res, "text LIKE ?", filter).Error
	}
	return res, err
}

func TweetByID(id string) (res request.Tweet, err error) {
	err = db.Preload("User").Preload("Geo").First(&res, request.Tweet{ID: id}).Error
	return res, err
}

func TweetsByUser(username string, n uint, startTime, endTime *time.Time) (res []request.Tweet, err error) {
	query := db.Joins("INNER JOIN users ON users.id = tweets.user_id AND users.username = ?", username).Limit(int(n)).Preload("User").Preload("Geo")
	if startTime != nil && endTime != nil {
		err = query.Find(&res, "tweets.created_at >= ? AND tweets.created_at <= ?", startTime, endTime).Error
	} else {
		err = query.Find(&res).Error
	}
	return res, err
}

func InsertPoliticians(politicians []request.Politician) (err error) {
	return db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&politicians).Error
}

func PoliticianByNameSurname(name string, surname string) (res request.Politician, err error) {
	err = db.First(&res, request.Politician{Name: strings.ToUpper(name), Surname: strings.ToUpper(surname)}).Error
	return res, err
}
func PoliticiansScoreboard() (res []request.Politician, err error) {
	err = db.Model(&request.Politician{}).Order("points desc").Find(&res).Error
	return res, err
}
func PoliticianBestSingleScore() (res request.Politician, err error) {
	err = db.Model(&request.Politician{}).Order("best_single_score desc").First(&res).Error
	return res, err
}
func PoliticianBestAverage() (res request.Politician, err error) {
	err = db.Model(&request.Politician{}).Order("average desc").First(&res).Error
	return res, err
}

func AddPointsPoliticianByNameSurname(p request.Politician) (err error) {
	politician, err := PoliticianByNameSurname(p.Name, p.Surname)
	if err != nil {
		err = InsertPoliticians([]request.Politician{p})
		if err != nil {
			return
		}
		politician, _ = PoliticianByNameSurname(p.Name, p.Surname)
	} else {
		// update just in case it's new
		if politician.LastUpdated.Before(p.LastUpdated) {
			politician.Points += p.Points
			politician.LastUpdated = p.LastUpdated
			return db.Model(&request.Politician{}).Where(request.Politician{ID: politician.ID}).Update("points", politician.Points).Update("last_updated", politician.LastUpdated).Error
		}
	}
	return
}

func AddPointsPoliticians(politicians []request.Politician) (err error) {
	for _, p := range politicians {
		if err = AddPointsPoliticianByNameSurname(p); err != nil {
			return
		}
	}
	return
}
