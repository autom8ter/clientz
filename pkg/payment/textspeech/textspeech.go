package textspeech

import (
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"context"
	"log"
)

func New(ctx context.Context) *texttospeech.Client {
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
