package goals

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmlopezz/uluru-api/internal/util"
	authclient "github.com/microapis/authentication-api/client"
)

type handlerContext struct {
	GoalStore           GoalStore
	RetirementGoalStore RetirementGoalStore
	AuthClient          *authclient.Client
}

// Routes ...
func Routes(r *mux.Router, ac *authclient.Client, gs GoalStore, rs RetirementGoalStore) {
	// define context
	ctx := &handlerContext{
		GoalStore:           gs,
		RetirementGoalStore: rs,
	}

	//
	// PUBLIC ROUTES
	//
	p := r.PathPrefix("/api/v1/goals").Subrouter()
	// GET /api/v1/goals
	p.HandleFunc("", listGoals(ctx)).Methods(http.MethodGet, http.MethodOptions)

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
	u.HandleFunc("/retirements", getLastRetirement(ctx)).Methods(http.MethodGet, http.MethodOptions)
	// POST /api/v1/goals/retirements
	u.HandleFunc("/retirements", createRetirementGoal(ctx)).Methods(http.MethodPost, http.MethodOptions)
}
