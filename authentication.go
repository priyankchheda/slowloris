package main

import (
	"net/http"

	"github.com/mrjones/oauth"
)

// twitter oauth http client
var client *http.Client

// GetClient twitter authentication function
// Returns a http.Client for further communication with Twitter API server
func GetClient(
	consumerKey string,
	consumerSecret string,
	accessToken string,
	accessTokenSecret string) {
	c := oauth.NewConsumer(
		consumerKey,
		consumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
			AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
			AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
		})

	t := oauth.AccessToken{
		Token:  accessToken,
		Secret: accessTokenSecret,
	}

	client, _ = c.MakeHttpClient(&t)
}
