package users

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

type handlerContext struct {
	UserService Service
}

// Routes ...
func Routes(r *mux.Router, us Service) {
	// define context
	ctx := &handlerContext{
		UserService: us,
	}

	u1 := r.PathPrefix("/api/v1/users").Subrouter()
	// GET /api/v1/users/
	u1.HandleFunc("", listUsers(ctx)).Methods(h.MethodGet, h.MethodOptions)
	// POST /api/v1/users
	u1.HandleFunc("", createUser(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// routes they use user_id pathparams
	u2 := r.PathPrefix("/api/v1/users").Subrouter()
	u2.Use(GetUserIDParam())
	// GET /api/v1/users/:id
	u2.HandleFunc("/{user_id}", getUser(ctx)).Methods(h.MethodGet, h.MethodOptions)
	// PUT /api/v1/users/:id
	u2.HandleFunc("/{user_id}", updateUser(ctx)).Methods(h.MethodPut, h.MethodOptions)
	// Delete /api/v1/users/:id
	u2.HandleFunc("/{user_id}", deleteUser(ctx)).Methods(h.MethodDelete, h.MethodOptions)
}
