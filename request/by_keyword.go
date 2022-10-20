package request

func TweetsByKeyword(keyword string, n uint) (tweets []Tweet, err error) {
	url, err := buildURL(NewRequest("tweets/search/recent").
		WithQuery(keyword).
		AddTweetFields(RequestFieldAuthorID, RequestFieldGeo).
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
		).AddExpansions(RequestExpansionAuthorID),
	)
	if err != nil {
		return
	}
	return requestTweets(url, n)
}
