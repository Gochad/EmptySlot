package stripeGateway

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v76"
)

func setupStripeKey() {
	_ = godotenv.Load()
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

func Setup(ctx context.Context) {
	setupStripeKey()
}
