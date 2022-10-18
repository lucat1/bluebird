package main

import (
	// "git.hjkl.gq/bluebird/bluebird/api"
	"git.hjkl.gq/bluebird/bluebird/request"
	// "log"
	// "net/http"
)

const ADDR = ":8080"

func main() {
	// http.HandleFunc("/hello", api.HelloWorld)
	// if err := http.ListenAndServe(ADDR, nil); err != nil {
	// 	log.Fatalf("Could not listen on %s: %e", ADDR, err)
	// }

	request.Init()
	request.Test()
}
