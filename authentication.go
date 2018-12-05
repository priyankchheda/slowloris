package main

import (
	"net/http"

	"github.com/mrjones/oauth"
)

var client *http.Client

// GetClient twitter authentication function
func GetClient(consumerKey, consumerSecret, accessToken, accessTokenSecret string) {
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
