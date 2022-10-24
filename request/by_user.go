package request

func TweetsByUser(username string, n uint) (tweets []Tweet, err error) {
	user, err := requestUserByUsername(username)
	// TODO: not really a solution just a temp fix
	if err != nil && username != "" {
		return
	}
	url, err := buildURL(NewRequest("users/"+user.ID+"/tweets").
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
