package clientz

import (
	"cloud.google.com/go/firestore"
	language "cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/storage"
	"firebase.google.com/go"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudkms/v1"
	"net/http"
)

type Commander struct {
	Project string
	Bucket  string
	Handler http.HandlerFunc
	store   *firestore.Client
	config  *firebase.Config
	blob    *storage.Client
	auth    *cloudkms.Service
	creds   *google.Credentials
	lang    *language.Client
}
