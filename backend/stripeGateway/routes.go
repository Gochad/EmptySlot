package stripeGateway

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"backend/internal/models"
)

type item struct {
	id string
}

func CreatePayment(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Items []item `json:"items"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewDecoder.Decode: %v", err)
		return
	}

	resv := models.Reservation{
		Merchandise: models.Merchandise{
			ID:          "3",
			Name:        "no siema siema",
			Price:       10,
			Description: "superancki itemek",
		},
	}

	pl, _ := GeneratePaymentLink(SendToStripe(resv))

	writeJSON(w, struct {
		PaymentLink string `json:"paymentLink"`
	}{
		PaymentLink: pl.URL,
	})
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("json.NewEncoder.Encode: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := io.Copy(w, &buf); err != nil {
		log.Printf("io.Copy: %v", err)
		return
	}
}
