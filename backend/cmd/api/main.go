package main

import (
	"emptyslot/cmd/server"
	"emptyslot/internal"
	"emptyslot/internal/database"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/product"
	"log"
	"os"
)

type ResponseData struct {
	Message string `json:"message"`
}

func main() {
	database.ConnectDb()
	r := mux.NewRouter()
	internal.InitRoutes(r)
	server.NewServer(r)

}

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

	cos := &stripe.PaymentIntentListParams{}
	cos.Filters.AddFilter("limit", "", "3")
	i := paymentintent.List(cos)
	for i.Next() {
		log.Println(i.PaymentIntent())
	}
}
