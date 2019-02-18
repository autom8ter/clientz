package kms

import (
	cloudkms "cloud.google.com/go/kms/apiv1"
	"context"
	"log"
)

func New() *cloudkms.KeyManagementClient {
	ctx := context.Background()
	client, err := cloudkms.NewKeyManagementClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
