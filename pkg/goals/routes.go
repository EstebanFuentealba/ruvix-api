package goals

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmlopezz/afp-simulator/simulator"
	"github.com/jmlopezz/uluru-api/internal/util"
	authclient "github.com/microapis/authentication-api/client"
)

type handlerContext struct {
	GoalStore GoalStore
	Simulator simulator.Afp
}

// Routes ...
func Routes(r *mux.Router, ac *authclient.Client, gs GoalStore) {
	// define context
	ctx := &handlerContext{
		GoalStore: gs,
		Simulator: simulator.New(),
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
	a.Use(util.ValidateJWTWithRole(ac, "admin"))
	// POST /api/v1/goals
	a.HandleFunc("", createGoal(ctx)).Methods(http.MethodPost, http.MethodOptions)

	//
	// USER ROUTES
	//
	u := r.PathPrefix("/api/v1/goals").Subrouter()
	u.Use(util.ValidateJWTWithRole(ac, "user"))
	// GET /api/v1/goals/retirements
	u.HandleFunc("/retirements/last", getLastRetirement(ctx)).Methods(http.MethodGet, http.MethodOptions)
	// POST /api/v1/goals/retirements
	u.HandleFunc("/retirements", createRetirementGoal(ctx)).Methods(http.MethodPost, http.MethodOptions)
}
