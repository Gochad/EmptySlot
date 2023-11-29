package main

import (
	h "emptyslot/internal/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"github.com/stripe/stripe-go/v76/price"
	"github.com/stripe/stripe-go/v76/product"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type ResponseData struct {
	Message string `json:"message"`
}

func serverConfig(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":8000", // Wybierz dowolny numer portu
		Handler: handler,
	}
}

func closeServer(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	err := server.Shutdown(nil)
	if err != nil {
		log.Fatalf("Server closing error: %v", err)
	}
}

func runServer(server *http.Server) {
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
}

func main() {
	r := mux.NewRouter()
	s := r.PathPrefix("/merchandises").Subrouter()
	//s.HandleFunc("/", h.MerchandisesHandler)

	s.HandleFunc("/{key}/", h.MerchandiseHandler)
	s.HandleFunc("/", h.MerchandisesHandler)

	//s.HandleFunc("/{key}/details", h.MerchandiseDetailsHandler)

	server := serverConfig(":8000", r)

	runServer(server)

	closeServer(server)
}

//	func main() {
//		//something()
//		p := database.NewPostgres()
//		fmt.Println(p)
//		http.HandleFunc("/", routes.ReturnSingleMerchandise)
//
//		fmt.Println("Server is listening on :8080...")
//		http.ListenAndServe(":8080", nil)
//	}
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
