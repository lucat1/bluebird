package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"git.hjkl.gq/team14/team14/cache"
	"git.hjkl.gq/team14/team14/request"
	"git.hjkl.gq/team14/team14/server"
	"github.com/joho/godotenv"
	"gorm.io/gorm/logger"
)

const ADDR = ":8080"

func main() {
	if godotenv.Load() != nil {
		log.Println("Missing environment variable")
	}

	rand.Seed(time.Now().Unix())
	// v2 request client creation
	bearer := os.Getenv("TWITTER_BEARER")
	if bearer == "" {
		log.Fatalln("Missing environment variable TWITTER_BEARER")
	}
	client, err := request.NewClient("https://api.twitter.com/2/", bearer)
	if err != nil {
		log.Fatalf("Could not create http.Client: %v", err)
	}
	request.SetClient(client)

	sentiment := os.Getenv("SENTIMENT_SERVER")
	if sentiment == "" {
		log.Println("Missing environment variable SENTIMENT_SERVER, the analysis will be random")
	}
	request.SetSentimentURL(sentiment)

	// v1 request client creation
	consumerKey := os.Getenv("CONSUMER_KEY")
	if consumerKey == "" {
		log.Fatalln("Missing environment variable CONSUMER_KEY")
	}
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	if consumerSecret == "" {
		log.Fatalln("Missing environment variable CONSUMER_SECRET")
	}
	oauthToken := os.Getenv("OAUTH_TOKEN")
	if oauthToken == "" {
		log.Fatalln("Missing environment variable OAUTH_TOKEN")
	}
	oauthSecret := os.Getenv("OAUTH_SECRET")
	if oauthSecret == "" {
		log.Fatalln("Missing environment variable OAUTH_SECRET")
	}
	v1Client, err := request.NewV1Client("https://upload.twitter.com/1.1/media/upload.json?media_category=tweet_image", "https://api.twitter.com/2/", consumerKey, consumerSecret, oauthToken, oauthSecret)
	if err != nil {
		log.Fatalf("Could not create http.Client (v1): %v", err)
	}
	request.SetV1Client(v1Client)

	if err = cache.Open("bluebird.db", logger.Warn); err != nil {
		log.Fatalf("Could not open database: %v", err)
	}
	tweets, err := cache.TweetsAll()
	if err != nil {
		log.Fatalf("Could not fetch all cached tweets: %v", err)
	}
	log.Printf("Tweets in cache: %d", len(tweets))
	defer cache.Close()

	go func() {
		startTime, _ := time.Parse(time.RFC3339, "2022-09-29T11:02:59.263Z")
		politicians, err := request.PoliticiansScore(2000, startTime, time.Now())
		if err == nil {
			cache.AddPointsPoliticians(politicians)
		}
		teams, err := request.Teams()
		if err == nil {
			cache.InsertTeams(teams)
		}
	}()

	if err = server.RunServer(ADDR); err != nil {
		log.Fatalf("Could not open HTTP server: %v", err)
	}

	log.Println("bye")
}
