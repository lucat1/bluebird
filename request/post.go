package request

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
)

var tweets, _ = url.Parse("tweets")

func sendPost(url string, bodyReq []byte, contentType string) (res []byte, err error) {
	resp, err := v1Client.HTTP.Post(url, contentType, bytes.NewBuffer(bodyReq))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	r, err := ioutil.ReadAll(resp.Body)
	return r, err
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
	rawRes, err := sendPost(v1Client.UploadURL.String(), body.Bytes(), mp.FormDataContentType())
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

	rawRes, err := sendPost(v1Client.APIURL.ResolveReference(tweets).String(), buf, "application/json")
	if err != nil {
		return
	}

	var rawTweetRes rawTweetResponse
	err = json.Unmarshal(rawRes, &rawTweetRes)
	return rawTweetRes.Data, err
}
