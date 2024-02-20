package stripeGateway

import (
	"fmt"

	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/product"
)

func Pay(name, description string, priceUnit int64) (string, error) {
	if priceUnit == 0 {
		return "", fmt.Errorf("price unit cannot be 0")
	}
	link, _ := generatePaymentLink(sendToStripe(name, description, priceUnit))
	if link != nil && link.Active {
		return link.URL, nil
	}

	return "", nil
}

func sendToStripe(name, description string, priceUnit int64) *stripe.Price {
	productParams := &stripe.ProductParams{
		Name:        stripe.String(name),
		Description: stripe.String(description),
	}

	starterProduct, _ := product.New(productParams)

	priceParams := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyPLN)),
		Product:  stripe.String(starterProduct.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		UnitAmount: stripe.Int64(priceUnit),
	}
	pi, _ := price.New(priceParams)

	return pi
}
