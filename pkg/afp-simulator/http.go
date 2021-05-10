package simulator

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
	s := r.PathPrefix("/api/v1/afp-simulator").Subrouter()

	// define context
	ctx := handlerContext{}

	// POST /api/v1/afpsimulator/estimation
	s.HandleFunc("/estimation", estimation(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// POST /api/v1/afpsimulator/balance
	s.HandleFunc("/balance", balance(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// POST /api/v1/afpsimulator/retirement
	s.HandleFunc("/retirement", retirement(ctx)).Methods(h.MethodPost, h.MethodOptions)
}
