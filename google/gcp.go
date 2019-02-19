package google

import (
	"context"
	"golang.org/x/oauth2/google"
	"log"
)

func New(ctx context.Context) *google.Credentials {
	creds, err := google.FindDefaultCredentials(ctx, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		log.Fatal(err)
	}
	return creds
}
