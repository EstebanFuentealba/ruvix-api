package savings

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
	s := r.PathPrefix("/api/v1/savings").Subrouter()

	// define context
	ctx := handlerContext{}

	// GET /api/v1/savings/institutions
	s.HandleFunc("/institutions", getInstitutions(ctx)).Methods(h.MethodGet, h.MethodOptions)

	// POST /api/v1/savings/institutions
	s.HandleFunc("/institutions", createInstitution(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// GET /api/v1/savings/accounts
	s.HandleFunc("/accounts", getAccounts(ctx)).Methods(h.MethodGet, h.MethodOptions)

	// POST /api/v1/savings/accounts
	s.HandleFunc("/accounts", cerateAccount(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// GET /api/v1/savings/instruments
	s.HandleFunc("/instruments", getInstruments(ctx)).Methods(h.MethodGet, h.MethodOptions)

	// POST /api/v1/savings/instruments
	s.HandleFunc("/instruments", crateInstrument(ctx)).Methods(h.MethodPost, h.MethodOptions)

}
