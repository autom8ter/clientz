package iot

import (
	"context"
	"golang.org/x/oauth2/google"
	cloudiot "google.golang.org/api/cloudiot/v1"
	"log"
)

func New() *cloudiot.Service {
	// Authorize the client using Application Default Credentials.
	// See https://g.co/dv/identity/protocols/application-default-credentials
	ctx := context.Background()
	httpClient, err := google.DefaultClient(ctx, cloudiot.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}
	client, err := cloudiot.New(httpClient)
	if err != nil {
		log.Fatal(err)
	}

	return client
}
