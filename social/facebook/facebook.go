package facebook

import (
	fb "github.com/huandu/facebook"
)

func New(appId, appSecret, redirect, signedReq string) *fb.App {
	// Create a global App var to hold app id and secret.
	globalApp := fb.New(appId, appSecret)

	return globalApp
}
