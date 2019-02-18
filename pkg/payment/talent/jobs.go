package talent

import (
	"context"
	"golang.org/x/oauth2/google"
	talent "google.golang.org/api/jobs/v3"
	"log"
)

func New() *talent.Service {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, talent.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	// Create the jobs service client.
	ctsService, err := talent.New(client)
	if err != nil {
		log.Fatal(err)
	}
	return ctsService
}
