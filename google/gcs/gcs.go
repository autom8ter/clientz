package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"log"
)

func New(ctx context.Context) *storage.Client {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
