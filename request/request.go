package request

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// url should never start with '/'
func requestRaw[T userResponse | tweetResponse](url *url.URL) (raw T, err error) {
	res, err := client.HTTP.Get(client.URL.ResolveReference(url).String())
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		return raw, fmt.Errorf("Non 200 status code")
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	return raw, json.Unmarshal(body, &raw)
}

func requestUser(url *url.URL) (user User, err error) {
	var raw userResponse
	if raw, err = requestRaw[userResponse](url); err != nil {
		return User{}, err
	}
	user = raw.User()
	if user.ID == "" {
		return user, fmt.Errorf("User not found")
	}
	return
}

func requestTweets(url *url.URL, n uint) (tweets []Tweet, err error) {
	var raw tweetResponse
	var twts []Tweet

	for uint(len(tweets)) < n {
		if raw, err = requestRaw[tweetResponse](url); err != nil {
			return
		}
		twts, err = raw.Tweets()
		if err != nil {
			return
		}
		needed := min(int(n)-len(tweets), len(twts))
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
