package routes

import (
	"context"
	"github.com/gorilla/mux"
)

func RegisterRoutes(ctx context.Context, r *mux.Router) {
	registerMerchandise(ctx, r)
}
