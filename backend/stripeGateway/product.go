package stripeGateway

import (
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/product"

	"backend/internal/models"
)

func sendToStripe(reservation models.Reservation) {
	productParams := &stripe.ProductParams{
		Name:        stripe.String(reservation.Merchandise.Name),
		Description: stripe.String(reservation.Merchandise.Description),
	}

	starterProduct, _ := product.New(productParams)

	priceParams := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyPLN)),
		Product:  stripe.String(starterProduct.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		UnitAmount: stripe.Int64(reservation.Merchandise.Price),
	}
	_, _ = price.New(priceParams)
}
