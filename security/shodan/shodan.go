package shodan

import (
	"gopkg.in/ns3777k/go-shodan.v3/shodan"
	"log"
)

func New() *shodan.Client {
	client := shodan.NewEnvClient(nil)
	if client == nil {
		log.Fatal("Failed to register Shodan client from environmental variables")
	}
	return client
}
