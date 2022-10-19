package request

import (
	"net/url"
)

type RequestField string

const (
	RequestFieldAuthorID        RequestField = "author_id"
	RequestFieldGeo                          = "geo"
	RequestFieldCreatedAt                    = "created_at"
	RequestFieldDescription                  = "description"
	RequestFieldEntities                     = "entities"
	RequestFieldID                           = "id"
	RequestFieldLocation                     = "location"
	RequestFieldName                         = "name"
	RequestFieldPinnedTweetID                = "pinned_tweet_id"
	RequestFieldProfileImageURL              = "profile_image_url"
	RequestFieldProtected                    = "protected"
	RequestFieldPublicMetrics                = "public_metrics"
	RequestFieldURL                          = "url"
	RequestFieldUsername                     = "username"
	RequestFieldVerified                     = "verified"
	RequestFieldWithheld                     = "withheld"
)

type RequestExpansions string

const (
	RequestExpansionAuthorID RequestExpansions = "author_id"
)

type RequestQuery string

const (
	RequestQueryQuery           RequestQuery = "query"
	RequestQueryTweetFields     RequestQuery = "tweet.fields"
	RequestQueryUserFields      RequestQuery = "user.fields"
	RequestQueryExpansions      RequestQuery = "expansions"
	RequestQueryPaginationToken RequestQuery = "pagination_token"
)

type RequestURL struct {
	base            string
	query           string
	tweetFields     []RequestField
	userFields      []RequestField
	expansions      []RequestExpansions
	paginationToken string
}

func NewRequest(base string) RequestURL {
	return RequestURL{
		base:            base,
		query:           "",
		tweetFields:     []RequestField{},
		userFields:      []RequestField{},
		expansions:      []RequestExpansions{},
		paginationToken: "",
	}
}

func (req RequestURL) WithQuery(query string) RequestURL {
	req.query = query
	return req
}

func (req RequestURL) AddTweetFields(fields ...RequestField) RequestURL {
	req.tweetFields = append(req.tweetFields, fields...)
	return req
}

func (req RequestURL) AddUserFields(fields ...RequestField) RequestURL {
	req.userFields = append(req.userFields, fields...)
	return req
}

func (req RequestURL) AddExpansions(expansions ...RequestExpansions) RequestURL {
	req.expansions = append(req.expansions, expansions...)
	return req
}

func join[K ~string](strs []K, sep string) (s string) {
	l := len(strs)
	for i, v := range strs {
		s += string(v)
		if i < l-1 {
			s += sep
		}
	}
	return
}

func queryAdd(query url.Values, k string, v string) {
	if v != "" {
		query.Add(k, v)
	}
}

func buildURL(req RequestURL) (parsed *url.URL, err error) {
	parsed, err = url.Parse(req.base)
	if err != nil {
		return
	}
	query := parsed.Query()
	queryAdd(query, string(RequestQueryQuery), req.query)
	queryAdd(query, string(RequestQueryTweetFields), join(req.tweetFields, ","))
	queryAdd(query, string(RequestQueryUserFields), join(req.userFields, ","))
	queryAdd(query, string(RequestQueryExpansions), join(req.expansions, ","))
	queryAdd(query, string(RequestQueryPaginationToken), req.paginationToken)
	parsed.RawQuery = query.Encode()
	return
}
