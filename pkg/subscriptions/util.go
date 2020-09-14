package subscriptions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jmlopezz/uluru-api"
	"github.com/microapis/transbank-sdk-golang"
	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

// TODO: the following values should read from ENV values
const (
	subscriptionCallbackURL           = "http://localhost:5000/api/v1/subscriptions/payment_callback"
	subscriptionCallbackURLProduction = "https://api.pensionatebien.cl/api/v1/subscriptions/payment_callback"
	subscriptionBillURL               = "http://localhost:5000/final/post/comprobante/webpay"
	subscriptionBillURLProduction     = "https://api.pensionatebien.cl/final/post/comprobante/webpay"
	webpayURL                         = "https://webpay3gint.transbank.cl/webpayserver/initTransaction?token_ws="
	webpayURLProduction               = "https://webpay3gint.transbank.cl/webpayserver/initTransaction?token_ws="
	redirectURL                       = "http://localhost:9000/subscription/payment"
	redirectURLProduction             = "https://www.pensionatebien.cl/subscription/payment"
)

// GetSubscriptionIDParam ...
func GetSubscriptionIDParam() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			vars := mux.Vars(r)
			subscriptionID := vars["subscription_id"]
			if subscriptionID == "" {
				err := "forbidden"
				fmt.Println(fmt.Sprintf("[Subscriptions][Error] %v", err))
				b, _ := json.Marshal(uluru.Response{Error: err})
				http.Error(w, string(b), http.StatusForbidden)
				return
			}

			// set subscriptions_id
			context.Set(r, "subscriptionID", subscriptionID)

			next.ServeHTTP(w, r)
		})
	}
}

// GetTransactionIDParam ...
func GetTransactionIDParam() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			vars := mux.Vars(r)
			transactionID := vars["transaction_id"]
			if transactionID == "" {
				err := "forbidden"
				fmt.Println(fmt.Sprintf("[Subscriptions][Error] %v", err))
				b, _ := json.Marshal(uluru.Response{Error: err})
				http.Error(w, string(b), http.StatusForbidden)
				return
			}

			// set transaction_id
			context.Set(r, "transactionID", transactionID)

			next.ServeHTTP(w, r)
		})
	}
}

// PaymentToken ...
func PaymentToken(t *Transaction) (string, error) {
	if t.Subscription.Price == 0 {
		return "", nil
	}
	if t.ProviderID == ProviderWebpayPlusNormal {
		return webpayPlusToken(t)
	}
	return "", nil
}

// PaymentURL ...
func PaymentURL(t *Transaction) string {
	var url string
	environment := os.Getenv("ENVIRONMENT")
	if environment == "production" {
		url = webpayURLProduction
	} else {
		url = webpayURL
	}

	if t.ProviderID == ProviderWebpayPlusNormal {
		return fmt.Sprintf("%s%s", url, t.PaymentToken)
	}

	return ""
}

// PaymentVerify ...
func PaymentVerify(t *Transaction) error {
	if t.ProviderID == ProviderWebpayPlusNormal {
		return webpayPlusVerify(t)
	}

	return nil
}

func webpayPlusToken(t *Transaction) (string, error) {
	var returnURL string
	var finalURL string

	environment := os.Getenv("ENVIRONMENT")
	if environment == "production" {
		returnURL = subscriptionCallbackURLProduction
		finalURL = subscriptionBillURLProduction
	} else {
		returnURL = subscriptionCallbackURL
		finalURL = subscriptionBillURL
	}

	amount := t.Subscription.Price * float64(t.Subscription.Months)
	sessionID := t.OrderNumber
	buyOrder := t.OrderNumber

	service := webpay.NewIntegrationPlusNormal()
	transaction, err := service.InitTransaction(transbank.InitTransaction{
		Amount:    amount,
		SessionID: sessionID,
		BuyOrder:  buyOrder,
		ReturnURL: returnURL,
		FinalURL:  finalURL,
	})
	if err != nil {
		return "", err
	}

	return transaction.Token, nil
}

func webpayPlusVerify(t *Transaction) error {
	service := webpay.NewIntegrationPlusNormal()
	result, err := service.GetTransactionResult(t.PaymentToken)
	if err != nil {
		return err
	}

	fmt.Println(result)

	code, err := strconv.Atoi(result.DetailOutput.ResponseCode)
	if err != nil {
		return err
	}

	if code != 0 {
		return fmt.Errorf("transaction rejected with code=%v", code)
	}

	return nil
}

// ValidProvider ...
func ValidProvider(provider string) bool {
	pp := []string{
		ProviderFree,
		ProviderWebpayPatpass,
		ProviderWebpayPlusNormal,
	}

	for _, v := range pp {
		if v == provider {
			return true
		}
	}

	return false
}

// GetPaymentRedirectURL ...
func GetPaymentRedirectURL() string {
	var url string

	environment := os.Getenv("ENVIRONMENT")
	if environment == "production" {
		url = redirectURLProduction
	} else {
		url = redirectURL
	}

	return url
}
