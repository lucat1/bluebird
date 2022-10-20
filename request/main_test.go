package request

import (
	"git.hjkl.gq/bluebird/bluebird/test"
	"os"
	"testing"
)

var (
	byKeywordResponse tweetResponse
	byKeywordClient   *RequestClient
)

func TestMain(m *testing.M) {
	srv := test.CreateServer(test.ReadFile("../mock/by_keyword.json"))
	defer srv.Close()
	var err error
	byKeywordClient, err = NewClient(srv.URL, "")
	if err != nil {
		panic(err)
	}
	test.ReadJSON("../mock/by_keyword.json", &byKeywordResponse)
	os.Exit(m.Run())
}
