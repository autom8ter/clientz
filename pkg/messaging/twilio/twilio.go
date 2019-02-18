package twilio

import (
	"github.com/sfreiberg/gotwilio"
)

func New() *gotwilio.Twilio {
	accountSid := "ABC123..........ABC123"
	authToken := "ABC123..........ABC123"
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)
	return twilio
}
