package clientz

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/language/apiv1"
	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/storage"
	"cloud.google.com/go/videointelligence/apiv1"
	"cloud.google.com/go/vision/apiv1"
	"firebase.google.com/go"
	"github.com/bluele/slack"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/huandu/facebook"
	"github.com/logpacker/PayPal-Go-SDK"
	"github.com/sfreiberg/gotwilio"
	stripe "github.com/stripe/stripe-go/client"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudiot/v1"
	"google.golang.org/api/cloudkms/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
	talent "google.golang.org/api/jobs/v3"
	docker "github.com/docker/docker/client"
)

type Google struct {
	Project string
	Bucket  string
	Handler http.HandlerFunc
	store   *firestore.Client
	config  *firebase.Config
	blob    *storage.Client
	auth    *cloudkms.Service
	creds   *google.Credentials
	lang    *language.Client
	pub 	*pubsub.Client
	spch 	*speech.Client
	vis 	*vision.ImageAnnotatorClient
	vid 	*videointelligence.Client
	robots 	*cloudiot.Service
	tlent 	*talent.Service
}

type Social struct {
	fb *facebook.App
	twit *twitter.Client
}

type Messenger struct {
	twil *gotwilio.Twilio
	slck *slack.Slack
}


type Payment struct {
	pal *paypalsdk.Client
	strp *stripe.API
}


type Containers struct {
	dkr *docker.Client
	kube *kubernetes.Clientset
}


