package plans

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
	s := r.PathPrefix("/api/v1/plans").Subrouter()

	// define context
	ctx := handlerContext{}

	// GET /api/v1/plans
	//
	// note: includes plan features
	s.HandleFunc("/", getPlans(ctx)).Methods(h.MethodGet, h.MethodOptions)

	// POST /api/v1/plans
	s.HandleFunc("/", cratePlan(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// GET /api/v1/plans/payment-providers
	s.HandleFunc("/payment-providers", getProviders(ctx)).Methods(h.MethodGet, h.MethodOptions)

	// POST /api/v1/plans/payment-providers
	s.HandleFunc("/payment-providers", createPaymentProvider(ctx)).Methods(h.MethodGet, h.MethodOptions)

	// GET /api/v1/plans/transactions
	s.HandleFunc("/transctions", getTransactions(ctx)).Methods(h.MethodGet, h.MethodOptions)

	// POST /api/v1/plans/transactions
	s.HandleFunc("/transctions", createTransactions(ctx)).Methods(h.MethodGet, h.MethodOptions)
}
