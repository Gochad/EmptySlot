package stripeGateway

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func payForReservation(amount int64) {
	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(2000),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		Currency: stripe.String(string(stripe.CurrencyPLN)),
	}
	_, _ = paymentintent.New(params)
}
