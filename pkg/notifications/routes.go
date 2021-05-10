package notifications

import (
	"github.com/gorilla/mux"
)

type handlerContext struct{}

// Routes ...
func Routes(r *mux.Router) {
	// s := r.PathPrefix("/api/v1/notifications").Subrouter()

	// // define context
	// ctx := handlerContext{}

	// GET /api/v1/notifications
	// s.HandleFunc("/", getPlans(ctx)).Methods(h.MethodGet, h.MethodOptions)

	// // PUT /api/v1/notifications/:id
	// s.HandleFunc("/:id", createNotification(ctx)).Methods(h.MethodPut, h.MethodOptions)
}
