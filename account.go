package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

// GetAccountSettings returns settings (including current trend, geo and sleep
// time information) for the authenticating user. This function is a wrapper
// over account/settings.
// API Doc Link: https://developer.twitter.com/en/docs/accounts-and-users/
// manage-account-settings/api-reference/get-account-settings
func GetAccountSettings() (string, error) {
	response, err := client.Get(
		"https://api.twitter.com/1.1/account/settings.json")
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

// GetAccountVerifyCredentialsParameter is a api parameter structure
// Please follow the Twitter API documentation for more details.
// API Doc Link: https://developer.twitter.com/en/docs/accounts-and-users/
// manage-account-settings/api-reference/get-account-verify_credentials
type GetAccountVerifyCredentialsParameter struct {
	IncludeEntities *string `json:"include_entities"`
	SkipStatus      *string `true json:"skip_status"`
	IncludeEmail    *string `json:"include_email"`
}

// GetAccountVerifyCredentials returns representation of the requesting user
// if authentication is successful; else returns AuthenticationInvalid error.
// Use this function to test if the supplied user credentials are valid.
// This funciton is a wrapper over account/verify_credentials
// API Doc Link: https://developer.twitter.com/en/docs/accounts-and-users/
// manage-account-settings/api-reference/get-account-verify_credentials
func GetAccountVerifyCredentials(data []byte) (string, error) {
	param := GetAccountVerifyCredentialsParameter{}
	err := json.Unmarshal([]byte(data), &param)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("GET",
		"https://api.twitter.com/1.1/account/verify_credentials.json", nil)
	if err != nil {
		return "", err
	}

	q := request.URL.Query()
	boolValidInput := [6]string{"true", "false", "t", "f", "1", "0"}

	if param.IncludeEntities != nil {
		for _, b := range boolValidInput {
			if *param.IncludeEntities == b {
				q.Add("include_entities", *param.IncludeEntities)
				break
			} else {
				return "", errors.New("Invalid Parameter: include_entities")
			}
		}
	}

	if param.SkipStatus != nil {
		for _, b := range boolValidInput {
			if *param.SkipStatus == b {
				q.Add("skip_status", *param.SkipStatus)
				break
			} else {
				return "", errors.New("Invalid Parameter: skip_status")
			}
		}
	}

	if param.IncludeEmail != nil {
		for _, b := range boolValidInput {
			if *param.IncludeEmail == b {
				q.Add("include_email", *param.IncludeEmail)
				break
			} else {
				return "", errors.New("Invalid Parameter: include_email")
			}
		}
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

// GetUsersProfileBannerParameter is a api parameter structure.
// Please follow the Twitter API documentation for more details
// API Doc Link: https://developer.twitter.com/en/docs/accounts-and-users/
// manage-account-settings/api-reference/get-users-profile_banner
type GetUsersProfileBannerParameter struct {
	UserID     *string `json:"user_id"`
	ScreenName *string `json:"screen_name"`
}

// GetUsersProfileBanner returns a map of the available size variations of the
// specified user's profile banner. If the user has not uploaded a profile
// banner, a HTTP 404 will be served instead.
// This function is a wrapper over users/profile_banner
// API Doc Link: https://developer.twitter.com/en/docs/accounts-and-users/
// manage-account-settings/api-reference/get-users-profile_banner
func GetUsersProfileBanner(data []byte) (string, error) {
	param := GetUsersProfileBannerParameter{}
	err := json.Unmarshal(data, &param)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("GET",
		"https://api.twitter.com/1.1/users/profile_banner.json", nil)
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

	request.URL.RawQuery = q.Encode()
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		return "{'data': 'user has not uploaded a profile banner'}", nil
	}

	bits, err := readResponse(response.Body)
	if err != nil {
		return "", err
	}

	return string(bits), nil
}
