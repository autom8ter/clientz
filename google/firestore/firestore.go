package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	"log"
	"os"
)

func New() *firestore.Client {
	ctx := context.Background()

	projectID := os.Getenv("GCLOUD_PROJECT")
	if projectID == "" {
		log.Fatalf("Set Firebase project ID via GCLOUD_PROJECT env variable.")
	}

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
