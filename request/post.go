package request

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"git.hjkl.gq/luca.tagliavini5/oauth1"
)

func Post(payload TweetRequest) (res TweetResponse, err error) {
	myurl, err := buildURL(NewRequest("tweets"))
	if err != nil {
		return
	}

	buf, err := json.Marshal(payload)
	if err != nil {
		return
	}
	rawRes, err := requestPostRaw[rawTweetResponse](client, myurl, bytes.NewBuffer(buf), "application/json")
	if err != nil {
		return
	}
	return rawRes.Data, nil
}

func sendPost(url string, bodyReq []byte, contentType string) (res []byte, err error) {
	config := &oauth1.Config{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		CallbackURL:    "http://team14.hjkl.gq/",
		Endpoint: oauth1.Endpoint{
			RequestTokenURL: "https://api.twitter.com/oauth/request_token",
			AuthorizeURL:    "https://api.twitter.com/oauth/authorize",
			AccessTokenURL:  "https://api.twitter.com/oauth/access_token",
		},
		Noncer: oauth1.Base64Noncer{},
	}
	clock := oauth1.NewFixedClock(time.Now())
	auther := &oauth1.Auther{
		Config: config,
		Clock:  &clock,
	}
	token := oauth1.NewToken(os.Getenv("OAUTH_TOKEN"), os.Getenv("OAUTH_SECRET"))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyReq))
	err = auther.SetRequestAuthHeader(req, token)
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func PostMedia(media []byte) (res MediaResponse, err error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	part, err := mp.CreateFormFile("media", "prova")
	_, err = io.Copy(part, bytes.NewReader(media))
	if err != nil {
		return
	}
	mp.Close()
	rawRes, err := sendPost("https://upload.twitter.com/1.1/media/upload.json?media_category=tweet_image", body.Bytes(), mp.FormDataContentType())
	if err != nil {
		return
	}
	err = json.Unmarshal(rawRes, &res)
	return
}

func PostCustom(payload TweetRequest) (res TweetResponse, err error) {
	buf, err := json.Marshal(payload)
	if err != nil {
		return
	}
	rawRes, err := sendPost("https://api.twitter.com/2/tweets", buf, "application/json")
	if err != nil {
		return
	}
	var rawTweetRes rawTweetResponse
	err = json.Unmarshal(rawRes, &rawTweetRes)
	return rawTweetRes.Data, err
}
