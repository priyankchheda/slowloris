package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// GetSettings get twitter account settings
func GetSettings(client *http.Client) string {
	response, err := client.Get(
		"https://api.twitter.com/1.1/account/settings.json")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	bits, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(bits)
}
