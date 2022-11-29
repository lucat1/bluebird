package request

import (
	"bytes"
	"encoding/json"
	"fmt"
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
func PostCustom(client *RequestClient, payload TweetRequest) (res TweetResponse, err error) {
	myurl, err := buildURL(NewRequest("tweets"))
	if err != nil {
		return
	}

	fmt.Println("myurl")
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
