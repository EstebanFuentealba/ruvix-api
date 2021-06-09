package subscriptions

import (
	"time"

	ruvixapi "github.com/cagodoy/ruvix-api"
)

// FactoryCreatePaySubscription ...
func FactoryCreatePaySubscription(addr string, opts ruvixapi.ClientOptions) (*Subscription, *Subscription, error) {
	before := &Subscription{
		Features: []*Feature{
			{
				Text: "example feature 1",
			},
			{
				Text: "example feature 2",
			},
			{
				Text: "example feature 3",
			},
		},
		Name:   "example subscription name",
		Price:  15,
		Months: 6,
	}

	sc, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := sc.CreateSubscription(before)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryCreateFreeSubscription ...
func FactoryCreateFreeSubscription(addr string, opts ruvixapi.ClientOptions) (*Subscription, *Subscription, error) {
	before := &Subscription{
		Features: []*Feature{
			{
				Text: "example feature 1",
			},
			{
				Text: "example feature 2",
			},
			{
				Text: "example feature 3",
			},
		},
		Name:   "example free subscription name",
		Price:  0,
		Months: 0,
	}

	sc, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := sc.CreateSubscription(before)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryListSubscriptions ...
func FactoryListSubscriptions(addr string, opts ruvixapi.ClientOptions) (*Subscription, []*Subscription, error) {
	_, before, err := FactoryCreatePaySubscription(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.ListSubscriptions()
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryListProviders ...
func FactoryListProviders(addr string, opts ruvixapi.ClientOptions) ([]*Provider, error) {
	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, err
	}

	after, err := gs.ListProviders()
	if err != nil {
		return nil, err
	}

	return after, nil
}

// FactoryFreeSubscribe ...
func FactoryFreeSubscribe(addr string, opts ruvixapi.ClientOptions) (*Subscription, *Transaction, error) {
	_, before, err := FactoryCreateFreeSubscription(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.Subscribe(before.ID, ProviderFree)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryPaySubscribe ...
func FactoryPaySubscribe(addr string, opts ruvixapi.ClientOptions) (*Subscription, *Transaction, error) {
	_, before, err := FactoryCreatePaySubscription(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.Subscribe(before.ID, ProviderWebpayPlusNormal)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryPayUnsubscribe ...
func FactoryPayUnsubscribe(addr string, opts ruvixapi.ClientOptions) (*Transaction, *Transaction, error) {
	beforeSubscription, before, err := FactoryPaySubscribe(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	time.Sleep(10 * time.Millisecond)

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.Unsubscribe(beforeSubscription.ID)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryFreeUnsubscribe ...
func FactoryFreeUnsubscribe(addr string, opts ruvixapi.ClientOptions) (*Transaction, *Transaction, error) {
	beforeSubscription, before, err := FactoryFreeSubscribe(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	time.Sleep(10 * time.Millisecond)

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.Unsubscribe(beforeSubscription.ID)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryPayRefresh ...
func FactoryPayRefresh(addr string, opts ruvixapi.ClientOptions) (*Transaction, *Transaction, error) {
	beforeSubscription, before, err := FactoryPaySubscribe(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.Refresh(beforeSubscription.ID, ProviderWebpayPlusNormal)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryFreeRefresh ...
func FactoryFreeRefresh(addr string, opts ruvixapi.ClientOptions) (*Transaction, *Transaction, error) {
	beforeSubscription, before, err := FactoryFreeSubscribe(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.Refresh(beforeSubscription.ID, ProviderWebpayPlusNormal)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryPayVerify ...
func FactoryPayVerify(addr string, opts ruvixapi.ClientOptions) (*Transaction, *Transaction, error) {
	beforeSubscription, before, err := FactoryPaySubscribe(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.Verify(beforeSubscription.ID)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryFreeVerify ...
func FactoryFreeVerify(addr string, opts ruvixapi.ClientOptions) (*Transaction, *Transaction, error) {
	beforeSubscription, before, err := FactoryFreeSubscribe(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.Verify(beforeSubscription.ID)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// ListTransactions ...
func ListTransactions(addr string, opts ruvixapi.ClientOptions) (*Transaction, []*Transaction, error) {
	_, before, err := FactoryPaySubscribe(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.ListTransactions()
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// LastTransaction ...
func LastTransaction(addr string, opts ruvixapi.ClientOptions) (*Transaction, *Transaction, error) {
	_, before, err := FactoryPaySubscribe(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.LastTransaction()
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}
