package profile

import (
	"net/http"

	"github.com/cagodoy/ruvix-api/internal/util"
	"github.com/gorilla/mux"

	"github.com/cagodoy/ruvix-api/pkg/auth"
)

type handlerContext struct {
	Store Store
}

// Routes ...
func Routes(r *mux.Router, as auth.Service, store Store) {
	// define context
	ctx := &handlerContext{
		Store: store,
	}

	//
	// USER ROUTES
	//
	s := r.PathPrefix("/api/v1/profile").Subrouter()
	// check if token is valid and get user_id
	s.Use(util.ValidateJWTWithRole(as, "user"))
	// POST /api/v1/profile
	s.HandleFunc("", createProfile(ctx)).Methods(http.MethodPost, http.MethodOptions)
	// PUT /api/v1/profile
	s.HandleFunc("", updateProfile(ctx)).Methods(http.MethodPut, http.MethodOptions)
	// GET /api/v1/profile
	s.HandleFunc("", getProfile(ctx)).Methods(http.MethodGet, http.MethodOptions)
}
