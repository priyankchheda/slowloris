package main

import (
	"encoding/json"
	"net/http"
)

// GetStatusesUserTimelineParameter is a api parameter structure
// Please follow the Twitter API documentation for more details
// API Doc Link: https://developer.twitter.com/en/docs/tweets/timelines/
// api-reference/get-statuses-user_timeline.html
type GetStatusesUserTimelineParameter struct {
	UserID         *string `json:"user_id"`
	ScreenName     *string `json:"screen_name"`
	SinceID        *string `json:"since_id"`
	Count          *string `json:"count"`
	MaxID          *string `json:"max_id"`
	TrimUser       *string `json:"trim_user"`
	ExcludeReplies *string `json:"exclude_replies"`
	IncludeRTS     *string `json:"include_rts"`
}

// GetStatusesUserTimeline returns a collection of the most recent Tweets
// posted by the user indicated by the screen_name or user_id parameters.
// This function is a wrapper over statuses/user_timeline.
// API Doc Link: https://developer.twitter.com/en/docs/tweets/timelines/
// api-reference/get-statuses-user_timeline.html
func GetStatusesUserTimeline(data []byte) (string, error) {
	param := GetStatusesUserTimelineParameter{}
	err := json.Unmarshal(data, &param)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("GET",
		"https://api.twitter.com/1.1/statuses/user_timeline.json", nil)
	if err != nil {
		return "", err
	}

	q := request.URL.Query()
	if param.UserID != nil {
		q.Add("user_id", *param.UserID)
	}

	if param.ScreenName != nil {
		q.Add("screen_name", *param.ScreenName)
	}

	if param.SinceID != nil {
		q.Add("since_id", *param.SinceID)
	}

	if param.Count != nil {
		q.Add("count", *param.Count)
	}

	if param.MaxID != nil {
		q.Add("max_id", *param.MaxID)
	}

	if param.TrimUser != nil {
		q.Add("trim_user", *param.TrimUser)
	}

	if param.ExcludeReplies != nil {
		q.Add("exclude_replies", *param.ExcludeReplies)
	}

	if param.IncludeRTS != nil {
		q.Add("include_rts", *param.IncludeRTS)
	}

	request.URL.RawQuery = q.Encode()
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bits, err := readResponse(response.Body)
	if err != nil {
		return "", err
	}

	return string(bits), nil
}
