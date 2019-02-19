package vision

import (
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"log"
)

func New(ctx context.Context) *vision.ImageAnnotatorClient {
	visionClient, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("vision.NewAnnotatorClient: %v", err)
	}
	return visionClient
}
