package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	ruvixapi "github.com/cagodoy/ruvix-api"
	auth "github.com/cagodoy/ruvix-api/pkg/auth"
	"github.com/gorilla/context"
)

// ValidateJWT ...
func ValidateJWT(as auth.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// get token from header
			token := r.Header.Get("Authorization")
			if token == "" {
				err := "forbidden"
				fmt.Println(fmt.Sprintf("[Auth][Error] %v", err))
				b, _ := json.Marshal(ruvixapi.Response{Error: err})
				http.Error(w, string(b), http.StatusForbidden)
				return
			}

			// check token if valid
			t, err := as.VerifyToken(token, auth.KindUser)
			if err != nil {
				err := "forbidden"
				fmt.Println(fmt.Sprintf("[Auth][Error] %v", err))
				b, _ := json.Marshal(ruvixapi.Response{Error: err})
				http.Error(w, string(b), http.StatusForbidden)
				return
			}

			// set token and user_id
			context.Set(r, "token", token)
			context.Set(r, "userID", t.UserID)

			next.ServeHTTP(w, r)
		})
	}
}

// ValidateJWTWithRole ...
func ValidateJWTWithRole(as auth.Service, role string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// get token from header
			token := r.Header.Get("Authorization")
			if token == "" {
				err := "forbidden"
				fmt.Println(fmt.Sprintf("[Auth][Error] %v", err))
				b, _ := json.Marshal(ruvixapi.Response{Error: err})
				http.Error(w, string(b), http.StatusForbidden)
				return
			}

			// check token if valid
			t, err := as.VerifyToken(token, auth.KindUser)
			if err != nil {
				err := "forbidden"
				fmt.Println(fmt.Sprintf("[Auth][Error] %v", err))
				b, _ := json.Marshal(ruvixapi.Response{Error: err})
				http.Error(w, string(b), http.StatusForbidden)
				return
			}
			// TODO(ca): check if token is valid for use with admin role

			// set token and user_id
			context.Set(r, "token", token)
			context.Set(r, "userID", t.UserID)

			next.ServeHTTP(w, r)
		})
	}
}
