package request

import "time"

func TweetsByKeyword(keyword string, n uint, startTime, endTime *time.Time) (tweets []Tweet, err error) {
	url, err := buildURL(NewRequest("tweets/search/recent").
		WithQuery(keyword).
		AddStartTime(startTime).
		AddEndTime(endTime).
		MaxResults(n).
		AddTweetFields(RequestFieldAuthorID, RequestFieldGeo, RequestFieldCreatedAt).
		AddUserFields(
			RequestFieldWithheld,
			RequestFieldVerified,
			RequestFieldUsername,
			RequestFieldURL,
			RequestFieldPublicMetrics,
			RequestFieldProtected,
			RequestFieldProfileImageURL,
			RequestFieldPinnedTweetID,
			RequestFieldName,
			RequestFieldLocation,
			RequestFieldID,
			RequestFieldEntities,
			RequestFieldDescription,
			RequestFieldCreatedAt,
		).
		AddPlaceFields(
			RequestFieldID,
			RequestFieldGeo,
			RequestFieldFullName,
		).
		AddExpansions(RequestExpansionAuthorID, RequestExpansionGeoPlaceID),
	)
	if err != nil {
		return
	}
	return requestTweets(url, n)
}
