package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	if consumerKey == "" {
		log.Fatal("CONSUMER_KEY env variable is not set")
	} else if consumerSecret == "" {
		log.Fatal("CONSUMER_SECRET env variable is not set")
	} else if accessToken == "" {
		log.Fatal("ACCESS_TOKEN env variable is not set")
	} else if accessTokenSecret == "" {
		log.Fatal("ACCESS_TOKEN_SECRET env variable is not set")
	}

	client, err := GetAuthenticatedClient(consumerKey, consumerSecret, accessToken, accessTokenSecret)
	if err != nil {
		log.Fatal(err)
	}
	data := GetSettings(client)
	fmt.Println("GetSettings return response:", data)
}
