package request

import (
	"net/url"
	"time"
)

type RequestSortOrder string

const (
	RequestSortOrderRecency   RequestSortOrder = "recency"
	RequestSortOrderRelevancy                  = "relevancy"
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
	RequestFieldFullName                     = "full_name"
	RequestFieldConversationID               = "conversation_id"
)

type RequestExpansions string

const (
	RequestExpansionAuthorID   RequestExpansions = "author_id"
	RequestExpansionGeoPlaceID                   = "geo.place_id"
)

type RequestQuery string

const (
	RequestQueryQuery           RequestQuery = "query"
	RequestQueryTweetIDs                     = "ids"
	RequestQuerySortOrder                    = "sort_order"
	RequestQueryTweetFields                  = "tweet.fields"
	RequestQueryUserFields                   = "user.fields"
	RequestQueryPlaceFields                  = "place.fields"
	RequestQueryStartTime                    = "start_time"
	RequestQueryEndTime                      = "end_time"
	RequestQueryExpansions                   = "expansions"
	RequestQueryPaginationToken              = "pagination_token"
)

type RequestQueryLang string
type RequestQueryConversationID string

const (
	RequestQueryLangIT RequestQueryLang = "it"
)

type RequestURL struct {
	base            string
	query           string
	ids             []string
	sortOrder       RequestSortOrder
	tweetFields     []RequestField
	userFields      []RequestField
	placeFields     []RequestField
	expansions      []RequestExpansions
	startTime       *time.Time
	endTime         *time.Time
	paginationToken string
}

func NewRequest(base string) RequestURL {
	return RequestURL{
		base:            base,
		query:           "",
		ids:             []string{},
		sortOrder:       "",
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
func (req RequestURL) Lang(lang RequestQueryLang) RequestURL {
	req.query += " lang:" + string(lang)
	return req
}
func (req RequestURL) ConversationID(conversationID RequestQueryConversationID) RequestURL {
	req.query += " conversation_id:" + string(conversationID)
	return req
}
func (req RequestURL) IDs(ids ...string) RequestURL {
	req.ids = append(req.ids, ids...)
	return req
}

func (req RequestURL) SortOrder(sort RequestSortOrder) RequestURL {
	req.sortOrder = sort
	return req
}

func (req RequestURL) AddTweetFields(fields ...RequestField) RequestURL {
	req.tweetFields = append(req.tweetFields, fields...)
	return req
}

func (req RequestURL) AddPlaceFields(fields ...RequestField) RequestURL {
	req.placeFields = append(req.placeFields, fields...)
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

func (req RequestURL) AddStartTime(startTime *time.Time) RequestURL {
	req.startTime = startTime
	return req
}

func (req RequestURL) AddEndTime(endTime *time.Time) RequestURL {
	req.endTime = endTime
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
	queryAdd(query, string(RequestQueryTweetIDs), join(req.ids, ","))
	queryAdd(query, string(RequestQuerySortOrder), string(req.sortOrder))
	queryAdd(query, string(RequestQueryTweetFields), join(req.tweetFields, ","))
	queryAdd(query, string(RequestQueryUserFields), join(req.userFields, ","))
	queryAdd(query, string(RequestQueryPlaceFields), join(req.placeFields, ","))
	queryAdd(query, string(RequestQueryExpansions), join(req.expansions, ","))
	if req.startTime != nil {
		queryAdd(query, string(RequestQueryStartTime), req.startTime.Format(time.RFC3339))
	}
	if req.endTime != nil {
		queryAdd(query, string(RequestQueryEndTime), req.endTime.Format(time.RFC3339))
	}
	queryAdd(query, string(RequestQueryPaginationToken), req.paginationToken)
	parsed.RawQuery = query.Encode()
	return
}
