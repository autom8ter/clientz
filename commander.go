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
	docker "github.com/docker/docker/client"
	"github.com/huandu/facebook"
	"github.com/logpacker/PayPal-Go-SDK"
	"github.com/sfreiberg/gotwilio"
	stripe "github.com/stripe/stripe-go/client"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudiot/v1"
	"google.golang.org/api/cloudkms/v1"
	talent "google.golang.org/api/jobs/v3"
	"k8s.io/client-go/kubernetes"
	firest "github.com/autom8ter/clientz/pkg/google/firestore"

)

type Google struct {
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
func NewGoogle() *Google {
	return &Google{
		store:  firest.New(),
		config: nil,
		blob:   nil,
		auth:   nil,
		creds:  nil,
		lang:   nil,
		pub:    nil,
		spch:   nil,
		vis:    nil,
		vid:    nil,
		robots: nil,
		tlent:  nil,
	}
}

type Social struct {
	fb *facebook.App
	twit *twitter.Client
}

func NewSocial() *Social {
	return &Social{
		fb:   nil,
		twit: nil,
	}
}

type Messenger struct {
	twil *gotwilio.Twilio
	slck *slack.Slack
}

func NewMessenger() *Messenger {
	return &Messenger{
		twil: nil,
		slck: nil,
	}
}

type Payment struct {
	pal *paypalsdk.Client
	strp *stripe.API
}

func NewPayment() *Payment {
	return &Payment{
		pal:  nil,
		strp: nil,
	}
}

type Containers struct {
	dkr *docker.Client
	kube *kubernetes.Clientset
}

func NewContainers() *Containers {
	return &Containers{
		dkr:  nil,
		kube: nil,
	}
}


