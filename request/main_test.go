package request

import (
	"git.hjkl.gq/bluebird/bluebird/test"
	"os"
	"testing"
)

var (
	byKeywordResponse tweetResponse
	byKeywordClient   *RequestClient
	byUserResponse    tweetResponse
	byUserClient      *RequestClient
)

func TestMain(m *testing.M) {
	var err error
	byKeywordServer := test.CreateServer(test.ReadFile("../mock/by_keyword.json"))
	defer byKeywordServer.Close()
	byUserServer := test.CreateServer(test.ReadFile("../mock/by_user.json"))
	defer byUserServer.Close()
	byKeywordClient, err = NewClient(byKeywordServer.URL, "")
	if err != nil {
		panic(err)
	}
	byUserClient, err = NewClient(byUserServer.URL, "")
	if err != nil {
		panic(err)
	}
	test.ReadJSON("../mock/by_keyword.json", &byKeywordResponse)
	test.ReadJSON("../mock/by_user.json", &byUserResponse)
	os.Exit(m.Run())
}
