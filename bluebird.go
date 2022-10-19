package main

import (
	"fmt"
	"git.hjkl.gq/bluebird/bluebird/request"
	"os"
)

const ADDR = ":8080"

func main() {
	bearer := os.Getenv("TWITTER_BEARER")
	client, err := request.NewClient("https://api.twitter.com/2/", bearer)
	if err != nil {
		panic(err)
	}
	request.SetClient(client)
	fmt.Println(request.TweetsByKeyword("#eredita", 10))
}
