package stripe

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/product"
)

func something() {
	err := godotenv.Load()
	log.Println(err)
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	product_params := &stripe.ProductParams{
		Name:        stripe.String("Starter Subscription"),
		Description: stripe.String("$12/Month subscription"),
	}
	starter_product, _ := product.New(product_params)

	price_params := &stripe.PriceParams{
		Currency: stripe.String(string(stripe.CurrencyPLN)),
		Product:  stripe.String(starter_product.ID),
		Recurring: &stripe.PriceRecurringParams{
			Interval: stripe.String(string(stripe.PriceRecurringIntervalMonth)),
		},
		UnitAmount: stripe.Int64(1),
	}
	starter_price, _ := price.New(price_params)

	params := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(2000),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
		Currency: stripe.String(string(stripe.CurrencyPLN)),
	}
	pe, _ := paymentintent.New(params)
	fmt.Println("ddd", pe)

	fmt.Println("Success! Here is your starter subscription product id: " + starter_product.ID)
	fmt.Println("Success! Here is your starter subscription price id: " + starter_price.ID)
}
