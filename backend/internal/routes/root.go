package routes

import (
	"context"

	"github.com/gorilla/mux"
)

func RegisterRoutes(ctx context.Context, r *mux.Router) {
	registerUser(ctx, r)
	registerCustomer(ctx, r)
	registerMerchandise(ctx, r)
	registerCategory(ctx, r)
	registerReservation(ctx, r)
}
