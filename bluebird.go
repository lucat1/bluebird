package main

import (
	"git.hjkl.gq/bluebird/bluebird/request"
	"os"
)

const ADDR = ":8080"

func main() {
	bearer := os.Getenv("TWITTER_BEARER")
	request.SetClient(request.NewClient(bearer))
	request.Test()
}
