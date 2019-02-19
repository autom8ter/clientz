package oauth

import "golang.org/x/oauth2"

//This List is a Work in Progress. Taken from https://github.com/golang/oauth2

var Facebook = oauth2.Endpoint{
	AuthURL:  "https://www.facebook.com/v3.1/dialog/oauth",
	TokenURL: "https://graph.facebook.com/v3.1/oauth/access_token",
}

// Endpoint is Foursquare's OAuth 2.0 endpoint.
var FourSquare = oauth2.Endpoint{
	AuthURL:  "https://foursquare.com/oauth2/authorize",
	TokenURL: "https://foursquare.com/oauth2/access_token",
}

// Endpoint is Github's OAuth 2.0 endpoint.
var Github = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}

// Endpoint is Instagram's OAuth 2.0 endpoint.
var Instagram = oauth2.Endpoint{
	AuthURL:  "https://api.instagram.com/oauth/authorize",
	TokenURL: "https://api.instagram.com/oauth/access_token",
}

// Endpoint is MailChimp's OAuth 2.0 endpoint.
// See http://developer.mailchimp.com/documentation/mailchimp/guides/how-to-use-oauth2/
var Mailchimp = oauth2.Endpoint{
	AuthURL:  "https://login.mailchimp.com/oauth2/authorize",
	TokenURL: "https://login.mailchimp.com/oauth2/token",
}

// Endpoint is Slack's OAuth 2.0 endpoint.
var Slack = oauth2.Endpoint{
	AuthURL:  "https://slack.com/oauth/authorize",
	TokenURL: "https://slack.com/api/oauth.access",
}

// Endpoint is Spotify's OAuth 2.0 endpoint.
var Spotify = oauth2.Endpoint{
	AuthURL:  "https://accounts.spotify.com/authorize",
	TokenURL: "https://accounts.spotify.com/api/token",
}

// Endpoint is Amazon's OAuth 2.0 endpoint.
var Amazon = oauth2.Endpoint{
	AuthURL:  "https://www.amazon.com/ap/oa",
	TokenURL: "https://api.amazon.com/auth/o2/token",
}

// Endpoint is PayPal's OAuth 2.0 endpoint in live (production) environment.
var Paypal = oauth2.Endpoint{
	AuthURL:  "https://www.paypal.com/webapps/auth/protocol/openidconnect/v1/authorize",
	TokenURL: "https://api.paypal.com/v1/identity/openidconnect/tokenservice",
}

// SandboxEndpoint is PayPal's OAuth 2.0 endpoint in sandbox (testing) environment.
var PaypalSandbox = oauth2.Endpoint{
	AuthURL:  "https://www.sandbox.paypal.com/webapps/auth/protocol/openidconnect/v1/authorize",
	TokenURL: "https://api.sandbox.paypal.com/v1/identity/openidconnect/tokenservice",
}

// Endpoint is Uber's OAuth 2.0 endpoint.
var Uber = oauth2.Endpoint{
	AuthURL:  "https://login.uber.com/oauth/v2/authorize",
	TokenURL: "https://login.uber.com/oauth/v2/token",
}
