package goals

import (
	"net/http"

	"github.com/cagodoy/ruvix-api/internal/util"
	"github.com/cagodoy/ruvix-api/pkg/auth"
	"github.com/gorilla/mux"
)

type handlerContext struct {
	GoalStore GoalStore
	// Simulator simulator.Afp
}

// Routes ...
func Routes(r *mux.Router, as auth.Service, gs GoalStore) {
	// define context
	ctx := &handlerContext{
		GoalStore: gs,
		// Simulator: simulator.New(),
	}

	//
	// PUBLIC ROUTES
	//
	p1 := r.PathPrefix("/api/v1/goals").Subrouter()
	// GET /api/v1/goals
	p1.HandleFunc("", listGoals(ctx)).Methods(http.MethodGet, http.MethodOptions)
	// POST /api/v1/goals/guest-retirements
	p1.HandleFunc("/guest-retirements", createGuestRetirementGoal(ctx)).Methods(http.MethodPost, http.MethodOptions)

	//
	// PUBLIC ROUTES
	//
	// - fingerprint
	//
	p2 := r.PathPrefix("/api/v1/goals").Subrouter()
	p2.Use(GetFingerprintParam())
	// GET /api/v1/goals
	p2.HandleFunc("", listGoals(ctx)).Methods(http.MethodGet, http.MethodOptions)
	// GET /api/v1/goals/guest-retirements
	p2.HandleFunc("/guest-retirements/{fingerprint}/last", getLastGuestRetirement(ctx)).Methods(http.MethodGet, http.MethodOptions)

	//
	// ADMIN ROUTES
	//
	a := r.PathPrefix("/api/v1/goals").Subrouter()
	a.Use(util.ValidateJWTWithRole(as, "admin"))
	// POST /api/v1/goals
	a.HandleFunc("", createGoal(ctx)).Methods(http.MethodPost, http.MethodOptions)

	//
	// USER ROUTES
	//
	u := r.PathPrefix("/api/v1/goals").Subrouter()
	u.Use(util.ValidateJWTWithRole(as, "user"))
	// GET /api/v1/goals/retirements
	u.HandleFunc("/retirements/last", getLastRetirement(ctx)).Methods(http.MethodGet, http.MethodOptions)
	// POST /api/v1/goals/retirements
	u.HandleFunc("/retirements", createRetirementGoal(ctx)).Methods(http.MethodPost, http.MethodOptions)
}
