package paypal

import (
	"github.com/logpacker/PayPal-Go-SDK"
	"log"
)

func New(id, secret string) *paypalsdk.Client {
	// Create a client instance
	c, err := paypalsdk.NewClient(id, secret, paypalsdk.APIBaseSandBox)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
