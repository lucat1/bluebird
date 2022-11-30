package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"git.hjkl.gq/team14/team14/cache"
	"git.hjkl.gq/team14/team14/chess"
	"git.hjkl.gq/team14/team14/request"
	"git.hjkl.gq/team14/team14/server"
	"github.com/jasonlvhit/gocron"
	"github.com/joho/godotenv"
	"gorm.io/gorm/logger"
)

const ADDR = ":8080"

func scheduler() {
	<-gocron.Start()
	log.Fatal("WARN The scheduler exited")
}

func main() {
	if godotenv.Load() != nil {
		log.Println("Missing environment variable")
	}

	rand.Seed(time.Now().Unix())
	bearer := os.Getenv("TWITTER_BEARER")
	if bearer == "" {
		log.Fatalln("Missing environment variable TWITTER_BEARER")
	}
	client, err := request.NewClient("https://api.twitter.com/2/", "https://upload.twitter.com/1.1/", bearer)
	if err != nil {
		log.Fatalf("Could not create http.Client: %v", err)
	}
	request.SetClient(client)

	sentiment := os.Getenv("SENTIMENT_SERVER")
	if sentiment == "" {
		log.Println("Missing environment variable SENTIMENT_SERVER, the analysis will be random")
	}
	request.SetSentimentURL(sentiment)

	if err = cache.Open("bluebird.db", logger.Warn); err != nil {
		log.Fatalf("Could not open database: %v", err)
	}
	tweets, err := cache.TweetsAll()
	if err != nil {
		log.Fatalf("Could not fetch all cached tweets: %v", err)
	}
	log.Printf("Tweets in cache: %d", len(tweets))
	defer cache.Close()
	err = chess.Resume()
	if err == nil {
		log.Printf("Resumed match state: %v", chess.GetMatch())
	}
	defer chess.Store()

	go scheduler()
	if err = server.RunServer(ADDR); err != nil {
		log.Fatalf("Could not open HTTP server: %v", err)
	}

	log.Println("bye")
}
