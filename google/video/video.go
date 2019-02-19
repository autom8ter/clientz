package video

import (
	video "cloud.google.com/go/videointelligence/apiv1"
	"context"
	"log"
)

func New(ctx context.Context) *video.Client {
	// Creates a client.
	client, err := video.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}
