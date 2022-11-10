package request

import (
	"os"
	"testing"

	"git.hjkl.gq/team14/team14/test"
	"github.com/stretchr/testify/assert"
)

var (
	byKeywordResponse tweetResponse
	byKeywordClient   *RequestClient
	byUserResponse    tweetResponse
	byUserClient      *RequestClient
)

func TestMain(m *testing.M) {
	var err error
	byKeywordServer := test.CreateMultiServer(map[string][]byte{
		"/tweets/search/recent": test.ReadFile("../mock/by_keyword.json"),
	})
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

func TestMin(t *testing.T) {
	assert.Equal(t, 2, min(2, 3), "The minimum between 2 and 3 is 2")
	assert.Equal(t, 2, min(3, 2), "The minimum between 2 and 3 is 2")
}
