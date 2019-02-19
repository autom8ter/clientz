package stripe

import (
	"github.com/stripe/stripe-go/client"
)

func New(token string) *client.API {
	cli := &client.API{}
	cli.Init(token, nil)
	return cli
}
