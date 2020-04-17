package notifications

import (
	h "net/http"

	"github.com/gorilla/mux"
)

// Response ...
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type handlerContext struct{}

// Routes ...
func Routes(r *mux.Router) {
	s := r.PathPrefix("/api/v1/notifications").Subrouter()

	// define context
	ctx := handlerContext{}

	// GET /api/v1/notifications
	s.HandleFunc("/", getPlans(ctx)).Methods(h.MethodGet, h.MethodOptions)

	// PUT /api/v1/notifications/:id
	s.HandleFunc("/:id", createNotification(ctx)).Methods(h.MethodPut, h.MethodOptions)
}
