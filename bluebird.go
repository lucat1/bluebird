package main

import (
	"fmt"
	"git.hjkl.gq/bluebird/bluebird/request"
	"os"
)

const ADDR = ":8080"

func main() {
	bearer := os.Getenv("TWITTER_BEARER")
	request.SetClient(request.NewClient(bearer))
	fmt.Println(request.TweetsByKeyword("#eredita", 10))
}
