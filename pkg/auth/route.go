package auth

import (
	h "net/http"

	"github.com/gorilla/mux"
)

// HTTPResponse ...
type HTTPResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type handlerContext struct {
	AuthService Service
}

// Routes ...
func Routes(r *mux.Router, ac Service) {
	s := r.PathPrefix("/api/v1/auth").Subrouter()

	// define context
	ctx := handlerContext{
		AuthService: ac,
	}

	// POST /api/v1/auth/login
	s.HandleFunc("/login", login(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// POST /api/v1/auth
	s.HandleFunc("/signup", signup(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// POST /api/v1/verify-email
	s.HandleFunc("/verify-email", verifyEmail(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// POST /api/v1/auth/logout
	s.HandleFunc("/logout", logout(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// POST /api/v1/auth/forgot-password
	s.HandleFunc("/forgot-password", forgotPassword(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// POST /api/v1/auth/recover-password
	s.HandleFunc("/recover-password", recoverPassword(ctx)).Methods(h.MethodPost, h.MethodOptions)

	// TODO: implement endpoint POST /api/v1/auth/resend-verification-email
}
