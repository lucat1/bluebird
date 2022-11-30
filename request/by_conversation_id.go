package request

import "fmt"

func TweetsByConversationID(conversationID string, n uint, startTime string, endTime string) (tweets []Tweet, err error) {
	url, err := buildURL(NewRequest("tweets/search/recent").
		ConversationID(RequestQueryConversationID(conversationID)).
		SortOrder(RequestSortOrderRecency).
		AddStartTime(RequestTime(startTime)).
		AddEndTime(RequestTime(endTime)).
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
	fmt.Println(url)
	if err != nil {
		return
	}
	return requestTweets(url, n)
}
