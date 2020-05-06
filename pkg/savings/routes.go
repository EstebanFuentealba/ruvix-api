package savings

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmlopezz/uluru-api/internal/util"
	authclient "github.com/microapis/authentication-api/client"
)

type handlerContext struct {
	InstitutionStore InstitutionStore
}

// Routes ...
func Routes(r *mux.Router, ac *authclient.Client, is InstitutionStore) {
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
	// USER ROUTES
	//
	u := r.PathPrefix("/api/v1/savings").Subrouter()
	u.Use(util.ValidateJWTWithRole(ac, "user"))
	// POST /api/v1/savings/institutions
	u.HandleFunc("/institutions", createInstitution(ctx)).Methods(http.MethodPost, http.MethodOptions)
}
