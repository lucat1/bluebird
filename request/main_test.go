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
	byKeywordServer := test.CreateSimpleServer(test.ReadFile("../mock/by_keyword.json"))
	defer byKeywordServer.Close()
	byUserServer := test.CreateMultiServer(map[string][]byte{
		"/users/by/username/salvinimi": test.ReadFile("../mock/id_by_user.json"),
		"/users/270839361/tweets":      test.ReadFile("../mock/by_user.json"),
	})
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
