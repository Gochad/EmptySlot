package stripegateway

import (
	"context"
	"os"

	"github.com/stripe/stripe-go/v76"
)

func setupStripeKey() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

func Setup(ctx context.Context) {
	setupStripeKey()
}
