package main

import (
	"fmt"
	"log"
	"os"
)

// Temp function
// As this project is supposed to be used as a package, will have to  delete
// this function in future
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

	GetClient(consumerKey, consumerSecret, accessToken, accessTokenSecret)

	data, err := GetAccountSettings()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nGetAccountSettings return response: \n\n", data)

	dataparam := `{"include_entities": "true", "include_email": "true"}`
	data, err = GetAccountVerifyCredentials([]byte(dataparam))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nGetAccountVerifyCredentials return response: \n\n", data)

	dataparam = `{"exclude_replies": "false"}`
	data, err = GetStatusesUserTimeline([]byte(dataparam))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nGetStatusesUserTimeline return response: \n\n", data)

	dataparam = `{}`
	data, err = GetUsersProfileBanner([]byte(dataparam))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\n\nGetUsersProfileBanner return response: \n\n", data)
}
