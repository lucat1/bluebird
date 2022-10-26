package main

import (
	"log"
	"os"

	"git.hjkl.gq/bluebird/bluebird/cache"
	"git.hjkl.gq/bluebird/bluebird/request"
	"git.hjkl.gq/bluebird/bluebird/server"
)

const ADDR = ":8080"

func main() {
	bearer := os.Getenv("TWITTER_BEARER")
	if bearer == "" {
		log.Fatalln("Missing environment variable TWITTER_BEARER")
	}
	client, err := request.NewClient("https://api.twitter.com/2/", bearer)
	if err != nil {
		log.Fatalf("Could not create http.Client: %v", err)
	}
	request.SetClient(client)
	if err = cache.Open(); err != nil {
		log.Fatalf("Could not open database: %v", err)
	}
	tweets, err := cache.TweetsAny()
	if err != nil {
		log.Fatalf("Could not fetch all cached tweets: %v", err)
	}
	log.Printf("Tweets in cache: %d", len(tweets))
	defer cache.Close()
	server.RunServer(ADDR)
}
