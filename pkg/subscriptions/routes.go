package subscriptions

import (
	"net/http"

	"github.com/cagodoy/ruvix-api/internal/util"
	"github.com/cagodoy/ruvix-api/pkg/auth"
	"github.com/gorilla/mux"
)

type handlerContext struct {
	SubscriptionStore SubscriptionStore
	AuthService       auth.Service
}

// Routes ...
func Routes(r *mux.Router, as auth.Service, ss SubscriptionStore) {
	// define context
	ctx := &handlerContext{
		SubscriptionStore: ss,
		AuthService:       as,
	}

	//
	// PUBLIC ROUTES
	//
	p := r.PathPrefix("/api/v1/subscriptions").Subrouter()
	// GET /api/v1/subscriptions
	p.HandleFunc("", listSubscriptions(ctx)).Methods(http.MethodGet, http.MethodOptions)
	// GET /api/v1/subscriptions/providers
	p.HandleFunc("/providers", listProviders(ctx)).Methods(http.MethodGet, http.MethodOptions)
	// GET /api/v1/subscriptions/providers
	p.HandleFunc("/payment_callback", paymentWebhook(ctx)).Methods(http.MethodPost, http.MethodOptions)

	//
	// ADMIN ROUTES
	//
	a := r.PathPrefix("/api/v1/subscriptions").Subrouter()
	a.Use(util.ValidateJWTWithRole(as, "admin"))
	// POST /api/v1/subscriptions
	a.HandleFunc("", createSubscription(ctx)).Methods(http.MethodPost, http.MethodOptions)

	//
	// USER ROUTES
	//
	u1 := r.PathPrefix("/api/v1/subscriptions").Subrouter()
	u1.Use(util.ValidateJWTWithRole(as, "user"))
	// GET /api/v1/subscriptions/transactions
	u1.HandleFunc("/transactions", listTransactions(ctx)).Methods(http.MethodGet, http.MethodOptions)
	// GET /api/v1/subscriptions/transactions/last
	u1.HandleFunc("/transactions/last", lastTransaction(ctx)).Methods(http.MethodGet, http.MethodOptions)

	//
	// USER ROUTES
	// - subscription_id
	//
	u2 := r.PathPrefix("/api/v1/subscriptions").Subrouter()
	u2.Use(util.ValidateJWTWithRole(as, "user"))
	u2.Use(GetSubscriptionIDParam())
	// POST /api/v1/subscriptions/:id/subscribe
	u2.HandleFunc("/{subscription_id}/subscribe", subscribe(ctx)).Methods(http.MethodPost, http.MethodOptions)
	// POST /api/v1/subscriptions/:id/unsubscribe
	u2.HandleFunc("/{subscription_id}/unsubscribe", unsubscribe(ctx)).Methods(http.MethodPost, http.MethodOptions)
	// POST /api/v1/subscriptions/:id/refresh
	u2.HandleFunc("/{subscription_id}/refresh", refresh(ctx)).Methods(http.MethodPost, http.MethodOptions)
	// POST /api/v1/subscriptions/:id/verify
	u2.HandleFunc("/{subscription_id}/verify", verify(ctx)).Methods(http.MethodPost, http.MethodOptions)
}
