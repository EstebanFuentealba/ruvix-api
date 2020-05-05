package subscriptions

import (
	"fmt"

	urulu "github.com/jmlopezz/uluru-api"
)

// FactoryCreateSubscription ...
func FactoryCreateSubscription(addr string, opts urulu.ClientOptions) (*Subscription, *Subscription, error) {
	before := &Subscription{
		Features: []*Feature{
			&Feature{
				Text: "example feature 1",
			},
			&Feature{
				Text: "example feature 2",
			},
			&Feature{
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

	fmt.Println(1111)
	after, err := sc.CreateSubscription(before)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(2222)

	return before, after, nil
}

// FactoryListSubscriptions ...
func FactoryListSubscriptions(addr string, opts urulu.ClientOptions) (*Subscription, []*Subscription, error) {
	_, before, err := FactoryCreateSubscription(addr, opts)
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
func FactoryListProviders(addr string, opts urulu.ClientOptions) ([]*Provider, error) {
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

// FactorySubscribe ...
func FactorySubscribe(addr string, opts urulu.ClientOptions) (*Subscription, *Transaction, error) {
	_, before, err := FactoryCreateSubscription(addr, opts)
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

// FactoryUnsubscribe ...
func FactoryUnsubscribe(addr string, opts urulu.ClientOptions) (*Transaction, *Transaction, error) {
	_, before, err := FactorySubscribe(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.Unsubscribe(before.ID)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// FactoryRefresh ...
func FactoryRefresh(addr string, opts urulu.ClientOptions) (*Transaction, *Transaction, error) {
	_, before, err := FactorySubscribe(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	gs, err := NewClient(addr, opts)
	if err != nil {
		return nil, nil, err
	}

	after, err := gs.Refresh(before.ID, ProviderWebpayPlusNormal)
	if err != nil {
		return nil, nil, err
	}

	return before, after, nil
}

// TODO(ca): implements FactoryVerify

// ListTransactions ...
func ListTransactions(addr string, opts urulu.ClientOptions) (*Transaction, []*Transaction, error) {
	_, before, err := FactorySubscribe(addr, opts)
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
func LastTransaction(addr string, opts urulu.ClientOptions) (*Transaction, *Transaction, error) {
	_, before, err := FactorySubscribe(addr, opts)
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
