package pub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"log"
	"os"
)

func New(ctx context.Context) *pubsub.Client {
	proj := os.Getenv("GOOGLE_CLOUD_PROJECT")
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Fatalf("Could not create pubsub Client: %v", err)
	}
}
