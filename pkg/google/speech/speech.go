package speech

import (
	"cloud.google.com/go/speech/apiv1"
	"context"
	"log"
)

func New(ctx context.Context) *speech.Client {
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
