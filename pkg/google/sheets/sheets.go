package sheets

import (
	"context"
	"golang.org/x/oauth2/google"
	spreadsheet "gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
	"log"
)

func New(ctx context.Context, credsPath string) *spreadsheet.Service {
	data, err := ioutil.ReadFile(credsPath)
	if err != nil {
		log.Fatal(err)
	}
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)
	return service
}
