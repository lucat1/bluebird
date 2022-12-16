package request

import (
	"os"
	"testing"

	"git.hjkl.gq/team14/team14/test"
	"github.com/stretchr/testify/assert"
)

var (
	byKeywordResponse                   tweetResponse
	byKeywordClient                     *RequestClient
	byUserResponse                      tweetResponse
	byUserClient                        *RequestClient
	rawSentimentResponse                sentimentResponse
	byKeywordGeoResponse                tweetResponse
	repliesClient                       *RequestClient
	byConvIDClient                      *RequestClient
	fantacitorioClient                  *RequestClient
	fantacitorioNoTweetsClient          *RequestClient
	ghigliottinaClient                  *RequestClient
	ghigliottinaTweetsErrorClient       *RequestClient
	ghigliottinaNoTweetsClient          *RequestClient
	ghigliottinaNoRepliesResponseClient *RequestClient
	ghigliottinaNoRepliesClient         *RequestClient
	ghigliottinaWinnersErrorClient      *RequestClient
	ghigliottinaTimeErrorClient         *RequestClient
)

func TestMain(m *testing.M) {
	var err error
	repliesServer := test.CreateMultiServer(map[string][]byte{
		"/tweets/search/recent": test.ReadFile("../mock/by_convid.json"),
		"/tweets":               test.ReadFile("../mock/by_tweetid.json"),
	})
	defer repliesServer.Close()
	byKeywordServer := test.CreateMultiServer(map[string][]byte{
		"/tweets/search/recent": test.ReadFile("../mock/by_keyword.json"),
	})
	defer byKeywordServer.Close()
	byUserServer := test.CreateMultiServer(map[string][]byte{
		"/users/by/username/salvinimi": test.ReadFile("../mock/id_by_user.json"),
		"/users/270839361/tweets":      test.ReadFile("../mock/by_user.json"),
	})
	defer byUserServer.Close()
	sentimentServer := test.CreateMultiServer(map[string][]byte{
		"/predict": test.ReadFile("../mock/sentiment.json"),
	})
	defer sentimentServer.Close()
	fantacitorioServer := test.CreateMultiServer(map[string][]byte{
		"/users/by/username/Fanta_citorio":  test.ReadFile("../mock/fantacitorio_user.json"),
		"/users/1492255549844566018/tweets": test.ReadFile("../mock/fantacitorio_tweets.json"),
	})
	defer fantacitorioServer.Close()

	fantacitorioNoTweetsServer := test.CreateMultiServer(map[string][]byte{
		"/users/by/username/Fanta_citorio": test.ReadFile("../mock/fantacitorio_user.json"),
	})
	defer fantacitorioNoTweetsServer.Close()

	ghigliottinaServer := test.CreateMultiServer(map[string][]byte{
		"/users/by/username/quizzettone":    test.ReadFile("../mock/ghigliottina_user.json"),
		"/users/1499992669480755204/tweets": test.ReadFile("../mock/ghigliottina_tweets.json"),
		"/tweets":                           test.ReadFile("../mock/ghigliottina_replies.json"),
		"/tweets/search/recent":             test.ReadFile("../mock/ghigliottina_conversation.json"),
	})
	defer ghigliottinaServer.Close()

	ghigliottinaTweetsErrorServer := test.CreateMultiServer(map[string][]byte{})
	defer ghigliottinaTweetsErrorServer.Close()

	ghigliottinaNoTweetsServer := test.CreateMultiServer(map[string][]byte{
		"/users/by/username/quizzettone":    test.ReadFile("../mock/ghigliottina_user.json"),
		"/users/1499992669480755204/tweets": test.ReadFile("../mock/ghigliottina_notweets.json"),
	})
	defer ghigliottinaNoTweetsServer.Close()

	ghigliottinaNoRepliesResponseServer := test.CreateMultiServer(map[string][]byte{
		"/users/by/username/quizzettone":    test.ReadFile("../mock/ghigliottina_user.json"),
		"/users/1499992669480755204/tweets": test.ReadFile("../mock/ghigliottina_tweets.json"),
	})
	defer ghigliottinaNoRepliesResponseServer.Close()

	ghigliottinaNoRepliesServer := test.CreateMultiServer(map[string][]byte{
		"/users/by/username/quizzettone":    test.ReadFile("../mock/ghigliottina_user.json"),
		"/users/1499992669480755204/tweets": test.ReadFile("../mock/ghigliottina_tweets.json"),
		"/tweets":                           test.ReadFile("../mock/ghigliottina_replies.json"),
		"/tweets/search/recent":             test.ReadFile("../mock/ghigliottina_conversation_noreplies.json"),
	})
	defer ghigliottinaNoRepliesServer.Close()

	ghigliottinaWinnersErrorServer := test.CreateMultiServer(map[string][]byte{
		"/users/by/username/quizzettone":    test.ReadFile("../mock/ghigliottina_user.json"),
		"/users/1499992669480755204/tweets": test.ReadFile("../mock/ghigliottina_tweets.json"),
		"/tweets":                           test.ReadFile("../mock/ghigliottina_replies.json"),
		"/tweets/search/recent":             test.ReadFile("../mock/ghigliottina_noconversation.json"),
	})
	defer ghigliottinaWinnersErrorServer.Close()

	ghigliottinaTimeErrorServer := test.CreateMultiServer(map[string][]byte{
		"/users/by/username/quizzettone":    test.ReadFile("../mock/ghigliottina_user.json"),
		"/users/1499992669480755204/tweets": test.ReadFile("../mock/ghigliottina_tweets.json"),
		"/tweets":                           test.ReadFile("../mock/ghigliottina_replies.json"),
		"/tweets/search/recent":             test.ReadFile("../mock/ghigliottina_conversation_timeerror.json"),
	})
	defer ghigliottinaTimeErrorServer.Close()

	repliesClient, err = NewClient(repliesServer.URL, "", "")
	if err != nil {
		panic(err)
	}
	byKeywordClient, err = NewClient(byKeywordServer.URL, "", "")
	if err != nil {
		panic(err)
	}
	byUserClient, err = NewClient(byUserServer.URL, "", "")
	if err != nil {
		panic(err)
	}
	fantacitorioClient, err = NewClient(fantacitorioServer.URL, "", "")
	if err != nil {
		panic(err)
	}
	fantacitorioNoTweetsClient, err = NewClient(fantacitorioNoTweetsServer.URL, "", "")
	if err != nil {
		panic(err)
	}
	ghigliottinaClient, err = NewClient(ghigliottinaServer.URL, "", "")
	if err != nil {
		panic(err)
	}
	ghigliottinaTweetsErrorClient, err = NewClient(ghigliottinaTweetsErrorServer.URL, "", "")
	if err != nil {
		panic(err)
	}
	ghigliottinaNoTweetsClient, err = NewClient(ghigliottinaNoTweetsServer.URL, "", "")
	if err != nil {
		panic(err)
	}
	ghigliottinaNoRepliesResponseClient, err = NewClient(ghigliottinaNoRepliesResponseServer.URL, "", "")
	if err != nil {
		panic(err)
	}
	ghigliottinaNoRepliesClient, err = NewClient(ghigliottinaNoRepliesServer.URL, "", "")
	if err != nil {
		panic(err)
	}
	ghigliottinaWinnersErrorClient, err = NewClient(ghigliottinaWinnersErrorServer.URL, "", "")
	if err != nil {
		panic(err)
	}
	ghigliottinaTimeErrorClient, err = NewClient(ghigliottinaTimeErrorServer.URL, "", "")
	if err != nil {
		panic(err)
	}

	SetSentimentURL(sentimentServer.URL)
	test.ReadJSON("../mock/by_keyword.json", &byKeywordResponse)
	test.ReadJSON("../mock/by_user.json", &byUserResponse)
	test.ReadJSON("../mock/sentiment.json", &rawSentimentResponse)
	test.ReadJSON("../mock/by_keyword_geo.json", &byKeywordGeoResponse)
	os.Exit(m.Run())
}

func TestMin(t *testing.T) {
	assert.Equal(t, 2, min(2, 3), "The minimum between 2 and 3 is 2")
	assert.Equal(t, 2, min(3, 2), "The minimum between 2 and 3 is 2")
}
