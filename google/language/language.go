package language

import (
	"cloud.google.com/go/language/apiv1"
	"context"
	"log"
)

func New(ctx context.Context) *language.Client {
	// Creates a client.
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}
