package stripegateway

import "os"

type config struct {
	TokenSecret string
}

var StripeConfig config

func init() {
	StripeConfig = config{
		TokenSecret: os.Getenv("STRIPE_SECRET_KEY"),
	}
}
