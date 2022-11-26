package request

import (
	"bytes"
	"encoding/json"
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
	rawRes, err := requestPostRaw[rawTweetResponse](client, myurl, bytes.NewBuffer(buf), "application/json", false)
	if err != nil {
		return
	}
	return rawRes.Data, nil
}
