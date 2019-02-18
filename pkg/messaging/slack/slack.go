package slack

import (
	"errors"
	"github.com/bluele/slack"
	"log"
)

func New(token string) {
	api := slack.New(token)
	if api == nil {
		log.Fatal(errors.New("Failed to register Slack client from token"))
	}
}
