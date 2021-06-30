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
		fmt.Println(fmt.Sprintf("[Subscriptions][List][Request] empty = %v", ""))

		subscriptions, err := ctx.SubscriptionStore.ListSubscriptions()
		if err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][List][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: subscriptions,
		}

		fmt.Println(fmt.Sprintf("[Subscriptions][List][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][List][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func listProviders(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[SubscriptionsProviders][List][Request] empty = %v", ""))

		providers, err := ctx.SubscriptionStore.ListProviders()
		if err != nil {
			fmt.Println(fmt.Sprintf("[SubscriptionsProviders][List][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: providers,
		}

		fmt.Println(fmt.Sprintf("[SubscriptionsProviders][List][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[SubscriptionsProviders][List][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func paymentWebhook(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Subscriptions][PaymentWebhook][Request] empty = %v", ""))

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][PaymentWebhook][Error] %v", err.Error()))
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
		fmt.Println(fmt.Sprintf("[Subscriptions][Create][Init]"))

		payload := &struct {
			Subscription *Subscription `json:"subscription"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Create][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Subscription == nil {
			err := "undefined subscription"
			fmt.Println(fmt.Sprintf("[Subscriptions][Create][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Subscription.Name == "" {
			err := "undefined name"
			fmt.Println(fmt.Sprintf("[Subscriptions][Create][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if len(payload.Subscription.Features) == 0 {
			err := "undefined features"
			fmt.Println(fmt.Sprintf("[Subscriptions][Create][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Println(fmt.Sprintf("[Subscriptions][Create][Request] payload = %v", payload))

		out, err := ctx.SubscriptionStore.CreateSubscription(payload.Subscription)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Create][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: out,
		}

		fmt.Println(fmt.Sprintf("[Subscriptions][Create][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Create][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func listTransactions(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[SubscriptionsTransactions][List][Request] empty = %v", ""))

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		transactions, err := ctx.SubscriptionStore.ListTransactions(QueryTransaction{
			UserID: userID,
		})
		if err != nil {
			fmt.Println(fmt.Sprintf("[SubscriptionsTransactions][List][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: transactions,
		}

		fmt.Println(fmt.Sprintf("[SubscriptionsTransactions][List][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[SubscriptionsTransactions][List][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func subscribe(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Init]"))

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		subscriptionID := context.Get(r, "subscriptionID").(string)
		if subscriptionID == "" {
			err := "subscriptionID is not defined"
			fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		payload := &struct {
			Subscribe *Transaction `json:"subscribe"`
		}{}

		if err := json.NewDecoder(r.Body).Decode(payload); err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Subscribe == nil {
			err := "undefined subscribe"
			fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Subscribe.ProviderID == "" && !ValidProvider(payload.Subscribe.ProviderID) {
			err := "undefined provider_id"
			fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Request] payload = %v", payload))

		// create transaction
		transaction, err := ctx.SubscriptionStore.Subscribe(QueryTransaction{
			UserID:         userID,
			SubscriptionID: subscriptionID,
			ProviderID:     payload.Subscribe.ProviderID,
		})
		if err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// get payment token
		token, err := PaymentToken(transaction)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		transaction.PaymentToken = token
		t, err := ctx.SubscriptionStore.UpdateTransaction(transaction)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Error] %v", err.Error()))
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

		fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Subscribe][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func unsubscribe(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Subscriptions][Unsubscribe][Init]"))

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Subscriptions][Unsubscribe][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		subscriptionID := context.Get(r, "subscriptionID").(string)
		if subscriptionID == "" {
			err := "subscriptionID is not defined"
			fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Error] %v", err))
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
			fmt.Println(fmt.Sprintf("[Subscriptions][Unsubscribe][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: data,
		}

		fmt.Println(fmt.Sprintf("[Subscriptions][Unsubscribe][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Unsubscribe][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func refresh(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Init]"))

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		subscriptionID := context.Get(r, "subscriptionID").(string)
		if subscriptionID == "" {
			err := "subscriptionID is not defined"
			fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Error] %v", err))
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
			fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Refresh == nil {
			err := "undefined refresh"
			fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if payload.Refresh.ProviderID == "" && ValidProvider(payload.Refresh.ProviderID) {
			err := "undefined provider_id"
			fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Request] payload = %v", payload))

		// refresh transaction
		transaction, err := ctx.SubscriptionStore.Refresh(QueryTransaction{
			UserID:         userID,
			SubscriptionID: subscriptionID,
			ProviderID:     payload.Refresh.ProviderID,
		})
		if err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		// get payment token
		token, err := PaymentToken(transaction)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		transaction.PaymentToken = token
		t, err := ctx.SubscriptionStore.UpdateTransaction(transaction)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Error] %v", err.Error()))
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

		fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Refresh][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func verify(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[Subscriptions][Verify][Request] empty = %v", ""))

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[Subscriptions][Verify][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		lastTransaction, err := ctx.SubscriptionStore.LastTransaction(userID)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Verify][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if lastTransaction.Status == StatusTransactionCompleted {
			err := errors.New("transaction status already complete")
			fmt.Println(fmt.Sprintf("[Subscriptions][Verify][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if lastTransaction.ProviderID == ProviderFree {
			err := fmt.Errorf("provider_id cannot be %s when subscription price is free", lastTransaction.ProviderID)
			fmt.Println(fmt.Sprintf("[Subscriptions][Verify][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		if lastTransaction.PaymentToken == "" {
			err := errors.New("undefined payment_token")
			fmt.Println(fmt.Sprintf("[Subscriptions][Verify][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		err = PaymentVerify(lastTransaction)
		if err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Verify][Error] %v", err.Error()))
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
			fmt.Println(fmt.Sprintf("[Subscriptions][Verify][Error7] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: transaction,
		}

		fmt.Println(fmt.Sprintf("[Subscriptions][Verify][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[Subscriptions][Verify][Error8] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}

func lastTransaction(ctx *handlerContext) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(fmt.Sprintf("[SubscriptionsTransactions][Last][Request] empty = %v", ""))

		userID := context.Get(r, "userID").(string)
		if userID == "" {
			err := "userID is not defined"
			fmt.Println(fmt.Sprintf("[SubscriptionsTransactions][Last][Error] %v", err))
			b, _ := json.Marshal(ruvixapi.Response{Error: err})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		transaction, err := ctx.SubscriptionStore.LastTransaction(userID)
		if err != nil {
			fmt.Println(fmt.Sprintf("[SubscriptionsTransactions][Last][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}

		res := ruvixapi.Response{
			Data: transaction,
		}

		fmt.Println(fmt.Sprintf("[SubscriptionsTransactions][Last][Response] %v", res))

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(fmt.Sprintf("[SubscriptionsTransactions][Last][Error] %v", err.Error()))
			b, _ := json.Marshal(ruvixapi.Response{Error: err.Error()})
			http.Error(w, string(b), http.StatusInternalServerError)
			return
		}
	}
}
