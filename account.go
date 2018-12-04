package main

import (
	"io/ioutil"
	"net/http"
)

// GetAccountSettings returns settings (including current trend, geo and sleep time information) for the authenticating user.
func GetAccountSettings(client *http.Client) (string, error) {
	response, err := client.Get("https://api.twitter.com/1.1/account/settings.json")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	err = CheckForResponseError(bits)
	if err != nil {
		return "", err
	}

	return string(bits), nil
}
