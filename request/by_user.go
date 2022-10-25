package request

func TweetsByUser(username string, n uint) (tweets []Tweet, err error) {
	url, err := buildURL(NewRequest("users/by/username/" + username))
	if err != nil {
		return
	}
	user, err := requestUser(url)
	if err != nil {
		return
	}
	url, err = buildURL(NewRequest("users/"+user.ID+"/tweets").
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
