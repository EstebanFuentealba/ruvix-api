package goals

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
	s := r.PathPrefix("/api/v1/goals").Subrouter()

	// define context
	ctx := handlerContext{}

	// GET /api/v1/goals
	//
	// note: get only de last of each type
	s.HandleFunc("/", getGoals(ctx)).Methods(h.MethodGet, h.MethodOptions)

	// POST /api/v1/goals
	s.HandleFunc("/", crateGoal(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// GET /api/v1/goals/retirements
	//
	// note: get only the last
	s.HandleFunc("/retirements", getRetirementGoals(ctx)).Methods(h.MethodGet, h.MethodOptions)

	// POST /api/v1/goals/retirements
	s.HandleFunc("/retirements", createRetirementGoal(ctx)).Methods(h.MethodPost, h.MethodOptions)
}
