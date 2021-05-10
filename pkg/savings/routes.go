package savings

import (
	"net/http"

	"github.com/cagodoy/ruvix-api/internal/util"
	"github.com/cagodoy/ruvix-api/pkg/auth"
	"github.com/gorilla/mux"
)

type handlerContext struct {
	InstitutionStore InstitutionStore
}

// Routes ...
func Routes(r *mux.Router, as auth.Service, is InstitutionStore) {
	// define context
	ctx := &handlerContext{
		InstitutionStore: is,
	}

	//
	// PUBLIC ROUTES
	//
	p := r.PathPrefix("/api/v1/savings").Subrouter()
	// GET /api/v1/savings/institutions
	p.HandleFunc("/institutions", listInstitutions(ctx)).Methods(http.MethodGet, http.MethodOptions)

	//
	// ADMIN ROUTES
	//
	u := r.PathPrefix("/api/v1/savings").Subrouter()
	u.Use(util.ValidateJWTWithRole(as, "admin"))
	// POST /api/v1/savings/institutions
	u.HandleFunc("/institutions", createInstitution(ctx)).Methods(http.MethodPost, http.MethodOptions)
}
