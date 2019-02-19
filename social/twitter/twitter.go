package twitter

import (
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"log"
)

func New(apptoken string) *twitter.Client {
	if apptoken == "" {
		log.Fatal("Please provide a valid twitter token")
	}
	config := &oauth2.Config{}
	token := &oauth2.Token{AccessToken: apptoken}
	// OAuth2 http.Client will automatically authorize Requests
	httpClient := config.Client(oauth2.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	return client
}
