package main

import (
	"net/http"
	"strconv"
)

// GetStatusesUserTimeline is a wrapper over statuses/user_timeline
// https://developer.twitter.com/en/docs/tweets/timelines/api-reference/get-statuses-user_timeline.html
func GetStatusesUserTimeline(
	userID string, screenName string, sinceID string, count int, maxID string,
	trimUser bool, excludeReplies bool, includeRTS bool) (string, error) {
	request, err := http.NewRequest("GET", "https://api.twitter.com/1.1/statuses/user_timeline.json", nil)
	q := request.URL.Query()
	if userID != "" {
		q.Add("user_id", userID)
	}
	if screenName != "" {
		q.Add("screen_name", screenName)
	}
	if sinceID != "" {
		q.Add("since_id", sinceID)
	}
	if maxID != "" {
		q.Add("max_id", maxID)
	}
	q.Add("count", strconv.Itoa(count))
	q.Add("trim_user", strconv.FormatBool(trimUser))
	q.Add("exclude_replies", strconv.FormatBool(excludeReplies))
	q.Add("include_rts", strconv.FormatBool(includeRTS))
	request.URL.RawQuery = q.Encode()

	response, err := client.Do(request)
	if err != nil {
		return "", nil
	}
	defer response.Body.Close()

	bits, err := readResponse(response.Body)
	if err != nil {
		return "", err
	}

	return string(bits), nil
}
