package goals

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jmlopezz/uluru-api"
)

// GetFingerprintParam ...
func GetFingerprintParam() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			vars := mux.Vars(r)
			fingerprint := vars["fingerprint"]
			if fingerprint == "" {
				err := "forbidden"
				fmt.Println(fmt.Sprintf("[Subscriptions][Error] %v", err))
				b, _ := json.Marshal(uluru.Response{Error: err})
				http.Error(w, string(b), http.StatusForbidden)
				return
			}

			// set subscriptions_id
			context.Set(r, "fingerprint", fingerprint)

			next.ServeHTTP(w, r)
		})
	}
}
