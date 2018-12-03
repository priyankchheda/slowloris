package main

import (
	"log"
	"net/http"

	"github.com/mrjones/oauth"
)

// GetAuthenticatedClient twitter authentication function
func GetAuthenticatedClient(consumerKey, consumerSecret, accessToken, accessTokenSecret string) (*http.Client, error) {
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessTokenSecret == "" {
		log.Fatal("Consumer key/secret and Access token/secret required")
	}

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

	return c.MakeHttpClient(&t)
}
