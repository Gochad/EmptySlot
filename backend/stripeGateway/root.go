package stripeGateway

import (
	"context"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v76"
)

func setupStripeKey() {
	_ = godotenv.Load()
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

func RegisterRoutes(ctx context.Context, r *mux.Router) {
	setupStripeKey()
	r.HandleFunc("/create-payment", CreatePayment).Methods("POST")
}
