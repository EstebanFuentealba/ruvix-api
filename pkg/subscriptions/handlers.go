package subscriptions

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	ruvixapi "github.com/cagodoy/ruvix-api"
	"github.com/gorilla/context"
)

func listSubscriptions(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[Subscriptions][List][Request] empty = %v\n", "")

		subscriptions, err := ctx.SubscriptionStore.ListSubscriptions()
		if err != nil {
			fmt.Printf("[Subscriptions][List][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: subscriptions,
		}

		fmt.Printf("[Subscriptions][List][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[Subscriptions][List][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func listProviders(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[SubscriptionsProviders][List][Request] empty = %v\n", "")

		providers, err := ctx.SubscriptionStore.ListProviders()
		if err != nil {
			fmt.Printf("[SubscriptionsProviders][List][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: providers,
		}

		fmt.Printf("[SubscriptionsProviders][List][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[SubscriptionsProviders][List][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func paymentWebhook(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[Subscriptions][PaymentWebhook][Request] empty = %v\n", "")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("[Subscriptions][PaymentWebhook][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
		log.Println(string(body))

		url := GetPaymentRedirectURL()
		http.Redirect(w, r, url, http.StatusMovedPermanently)

		fmt.Println("[Subscriptions][PaymentWebhook][Response]")
	}
}

func createSubscription(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[Subscriptions][Create][Init]")

		payload := &struct {
			Subscription *Subscription `json:"subscription"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Printf("[Subscriptions][Create][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Subscription == nil {
			err := "undefined subscription"
			fmt.Printf("[Subscriptions][Create][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Subscription.Name == "" {
			err := "undefined name"
			fmt.Printf("[Subscriptions][Create][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if len(payload.Subscription.Features) == 0 {
			err := "undefined features"
			fmt.Printf("[Subscriptions][Create][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Printf("[Subscriptions][Create][Request] payload = %v\n", payload)

		out, err := ctx.SubscriptionStore.CreateSubscription(payload.Subscription)
		if err != nil {
			fmt.Printf("[Subscriptions][Create][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: out,
		}

		fmt.Printf("[Subscriptions][Create][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[Subscriptions][Create][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func listTransactions(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[SubscriptionsTransactions][List][Request] empty = %v\n", "")

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Printf("[Subscriptions][Subscribe][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		transactions, err := ctx.SubscriptionStore.ListTransactions(QueryTransaction{
			UserID: userID,
		})
		if err != nil {
			fmt.Printf("[SubscriptionsTransactions][List][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: transactions,
		}

		fmt.Printf("[SubscriptionsTransactions][List][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[SubscriptionsTransactions][List][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func subscribe(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[Subscriptions][Subscribe][Init]")

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Printf("[Subscriptions][Subscribe][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		subscriptionID := context.Get(r, "subscriptionID").(string)
		if subscriptionID == "" {
			err := "subscriptionID is not defined"
			fmt.Printf("[Subscriptions][Subscribe][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		payload := &struct {
			Subscribe *Transaction `json:"subscribe"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Printf("[Subscriptions][Subscribe][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Subscribe == nil {
			err := "undefined subscribe"
			fmt.Printf("[Subscriptions][Subscribe][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Subscribe.ProviderID == "" && !ValidProvider(payload.Subscribe.ProviderID) {
			err := "undefined provider_id"
			fmt.Printf("[Subscriptions][Subscribe][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Printf("[Subscriptions][Subscribe][Request] payload = %v\n", payload)

		// create transaction
		transaction, err := ctx.SubscriptionStore.Subscribe(QueryTransaction{
			UserID:         userID,
			SubscriptionID: subscriptionID,
			ProviderID:     payload.Subscribe.ProviderID,
		})
		if err != nil {
			fmt.Printf("[Subscriptions][Subscribe][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// get payment token
		token, err := PaymentToken(transaction)
		if err != nil {
			fmt.Printf("[Subscriptions][Subscribe][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		transaction.PaymentToken = token
		t, err := ctx.SubscriptionStore.UpdateTransaction(transaction)
		if err != nil {
			fmt.Printf("[Subscriptions][Subscribe][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// prepare response
		res := ruvixapi.Response{
			Data: t,
		}

		if t.PaymentToken != "" {
			res.Meta = &TransactionMeta{
				PaymentURL: PaymentURL(t),
			}
		}

		fmt.Printf("[Subscriptions][Subscribe][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[Subscriptions][Subscribe][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func unsubscribe(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[Subscriptions][Unsubscribe][Init]")

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Printf("[Subscriptions][Unsubscribe][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		subscriptionID := context.Get(r, "subscriptionID").(string)
		if subscriptionID == "" {
			err := "subscriptionID is not defined"
			fmt.Printf("[Subscriptions][Refresh][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// update transaction
		data, err := ctx.SubscriptionStore.Unsubscribe(QueryTransaction{
			UserID:         userID,
			SubscriptionID: subscriptionID,
		})
		if err != nil {
			fmt.Printf("[Subscriptions][Unsubscribe][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: data,
		}

		fmt.Printf("[Subscriptions][Unsubscribe][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[Subscriptions][Unsubscribe][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func refresh(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[Subscriptions][Refresh][Init]")

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Printf("[Subscriptions][Refresh][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		subscriptionID := context.Get(r, "subscriptionID").(string)
		if subscriptionID == "" {
			err := "subscriptionID is not defined"
			fmt.Printf("[Subscriptions][Refresh][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		payload := &struct {
			Refresh *struct {
				ProviderID string `json:"provider_id"`
			} `json:"refresh"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Printf("[Subscriptions][Refresh][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Refresh == nil {
			err := "undefined refresh"
			fmt.Printf("[Subscriptions][Refresh][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Refresh.ProviderID == "" && ValidProvider(payload.Refresh.ProviderID) {
			err := "undefined provider_id"
			fmt.Printf("[Subscriptions][Refresh][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Printf("[Subscriptions][Refresh][Request] payload = %v\n", payload)

		// refresh transaction
		transaction, err := ctx.SubscriptionStore.Refresh(QueryTransaction{
			UserID:         userID,
			SubscriptionID: subscriptionID,
			ProviderID:     payload.Refresh.ProviderID,
		})
		if err != nil {
			fmt.Printf("[Subscriptions][Refresh][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// get payment token
		token, err := PaymentToken(transaction)
		if err != nil {
			fmt.Printf("[Subscriptions][Refresh][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		transaction.PaymentToken = token
		t, err := ctx.SubscriptionStore.UpdateTransaction(transaction)
		if err != nil {
			fmt.Printf("[Subscriptions][Refresh][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// prepare response
		res := ruvixapi.Response{
			Data: t,
		}

		if t.PaymentToken != "" {
			res.Meta = &TransactionMeta{
				PaymentURL: PaymentURL(t),
			}
		}

		fmt.Printf("[Subscriptions][Refresh][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[Subscriptions][Refresh][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func verify(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[Subscriptions][Verify][Request] empty = %v\n", "")

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Printf("[Subscriptions][Verify][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		lastTransaction, err := ctx.SubscriptionStore.LastTransaction(userID)
		if err != nil {
			fmt.Printf("[Subscriptions][Verify][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if lastTransaction.Status == StatusTransactionCompleted {
			err := errors.New("transaction status already complete")
			fmt.Printf("[Subscriptions][Verify][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if lastTransaction.ProviderID == ProviderFree {
			err := fmt.Errorf("provider_id cannot be %s when subscription price is free", lastTransaction.ProviderID)
			fmt.Printf("[Subscriptions][Verify][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if lastTransaction.PaymentToken == "" {
			err := errors.New("undefined payment_token")
			fmt.Printf("[Subscriptions][Verify][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		err = PaymentVerify(lastTransaction)
		if err != nil {
			fmt.Printf("[Subscriptions][Verify][Error] %v\n", err.Error())
			if strings.Contains(err.Error(), "rejected") {
				lastTransaction.Status = StatusTransactionRejected
			} else {
				b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
				http.Error(w, string(b), http.StatusInternalServerError)
				return
			}
		} else {
			lastTransaction.Status = StatusTransactionCompleted
		}

		transaction, err := ctx.SubscriptionStore.UpdateTransaction(lastTransaction)
		if err != nil {
			fmt.Printf("[Subscriptions][Verify][Error7] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: transaction,
		}

		fmt.Printf("[Subscriptions][Verify][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[Subscriptions][Verify][Error8] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func lastTransaction(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[SubscriptionsTransactions][Last][Request] empty = %v\n", "")

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Printf("[SubscriptionsTransactions][Last][Error] %v\n", err)
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		transaction, err := ctx.SubscriptionStore.LastTransaction(userID)
		if err != nil {
			fmt.Printf("[SubscriptionsTransactions][Last][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: transaction,
		}

		fmt.Printf("[SubscriptionsTransactions][Last][Response] %v\n", res)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Printf("[SubscriptionsTransactions][Last][Error] %v\n", err.Error())
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}
