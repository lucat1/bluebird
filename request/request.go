package request

import (
	"encoding/json"
	"io"
	"net/url"
)

func requestRaw(url *url.URL) (raw tweetResponse, err error) {
	res, err := client.HTTP.Get(client.URL.ResolveReference(url).String())
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	return raw, json.Unmarshal(body, &raw)
}

func requestTweets(url *url.URL, n uint) (tweets []Tweet, err error) {
	var raw tweetResponse
	var twts []Tweet

	for uint(len(tweets)) < n {
		if raw, err = requestRaw(url); err != nil {
			return
		}
		twts, err = raw.Tweets()
		if err != nil {
			return
		}
		needed := n - uint(len(tweets))
		tweets = append(tweets, twts[:needed]...)
		if raw.Meta.ResultCount == 0 || raw.Meta.NextToken == "" {
			break
		}
		queries := url.Query()
		queries.Set(string(RequestQueryPaginationToken), raw.Meta.NextToken)
		url.RawQuery = queries.Encode()
	}
	return
}
