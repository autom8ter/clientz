package docker

import (
	"github.com/docker/docker/client"
	"log"
)

func New() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err)
	}
	return cli
}
