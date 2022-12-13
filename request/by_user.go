package request

import "time"

func TweetsByUser(username string, n uint, startTime, endTime *time.Time) (tweets []Tweet, err error) {
	url, err := buildURL(NewRequest("users/by/username/" + username))
	if err != nil {
		return
	}
	user, err := requestUser(url)
	if err != nil {
		return
	}
	url, err = buildURL(NewRequest("users/"+user.ID+"/tweets").
		MaxResults(n).
		AddTweetFields(RequestFieldAuthorID, RequestFieldGeo, RequestFieldCreatedAt, RequestFieldEntities, RequestFieldAttachments).
		AddStartTime(startTime).
		AddEndTime(endTime).
		AddUserFields(
			RequestFieldWithheld,
			RequestFieldUsername,
			RequestFieldURL,
			RequestFieldProfileImageURL,
			RequestFieldName,
			RequestFieldLocation,
			RequestFieldID,
			RequestFieldCreatedAt,
		).AddMediaFields(RequestFieldURL).AddExpansions(RequestExpansionAuthorID, RequestExpansionAttachmentsMediaKeys),
	)
	if err != nil {
		return
	}
	return requestTweets(url, n)
}
