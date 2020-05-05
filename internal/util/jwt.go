package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/jmlopezz/uluru-api"
	auth "github.com/microapis/authentication-api"
	authclient "github.com/microapis/authentication-api/client"
)

// ValidateJWT ...
func ValidateJWT(ac *authclient.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// get token from header
			token := r.Header.Get("Authorization")
			if token == "" {
				err := "forbidden"
				fmt.Println(fmt.Sprintf("[Auth][Error] %v", err))
				b, _ := json.Marshal(uluru.Response{Error: err})
				http.Error(w, string(b), http.StatusForbidden)
				return
			}

			// check token if valid
			t, err := ac.VerifyToken(token, auth.KindUser)
			if err != nil {
				err := "forbidden"
				fmt.Println(fmt.Sprintf("[Auth][Error] %v", err))
				b, _ := json.Marshal(uluru.Response{Error: err})
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
func ValidateJWTWithRole(ac *authclient.Client, role string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			// get token from header
			token := r.Header.Get("Authorization")
			if token == "" {
				err := "forbidden"
				fmt.Println(fmt.Sprintf("[Auth][Error] %v", err))
				b, _ := json.Marshal(uluru.Response{Error: err})
				http.Error(w, string(b), http.StatusForbidden)
				return
			}

			// check token if valid
			t, err := ac.VerifyToken(token, auth.KindUser)
			if err != nil {
				err := "forbidden"
				fmt.Println(fmt.Sprintf("[Auth][Error] %v", err))
				b, _ := json.Marshal(uluru.Response{Error: err})
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
