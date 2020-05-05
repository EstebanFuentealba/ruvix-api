package subscriptions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	urulu "github.com/jmlopezz/uluru-api"
)

const (
	baseURLProduction  = "https://api.pensionatebien.cl/api/v1"
	baseURLStaging     = "https://api-staging.pensionatebien.cl/api/v1"
	baseURLDevelopment = "http://localhost:5000/api/v1"
)

// Client ...
type Client struct {
	client http.Client
	URL    string
	token  string
}

// NewClient ...
func NewClient(address string, opts urulu.ClientOptions) (*Client, error) {
	// TODO(ca): check on authorization if token is valid
	t := opts.Token

	c := http.Client{
		Timeout: time.Duration(50 * time.Second),
	}

	var url string
	if opts.Environment == "production" {
		url = baseURLProduction
	} else if opts.Environment == "staging" {
		url = baseURLStaging
	} else if opts.Environment == "development" {
		url = baseURLDevelopment
	} else {
		url = baseURLDevelopment
	}

	return &Client{
		client: c,
		URL:    url,
		token:  t,
	}, nil
}

// ListSubscriptions ...
func (c *Client) ListSubscriptions() ([]*Subscription, error) {
	url := fmt.Sprintf("%s/subscriptions", c.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	rr := struct {
		Data  []*Subscription `json:"data"`
		Error *string         `json:"error"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&rr); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if rr.Error != nil {
		return nil, errors.New(*rr.Error)
	}

	return rr.Data, nil
}

// ListProviders ...
func (c *Client) ListProviders() ([]*Provider, error) {
	url := fmt.Sprintf("%s/subscriptions/providers", c.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	rr := struct {
		Data  []*Provider `json:"data"`
		Error *string     `json:"error"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&rr); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if rr.Error != nil {
		return nil, errors.New(*rr.Error)
	}

	return rr.Data, nil
}

// CreateSubscription ...
func (c *Client) CreateSubscription(s *Subscription) (*Subscription, error) {
	r := struct {
		Subscription *Subscription `json:"subscription"`
	}{
		Subscription: s,
	}

	b, err := json.Marshal(&r)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/subscriptions", c.URL)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if c.token != "" {
		req.Header.Add("Authorization", c.token)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	rr := struct {
		Data  *Subscription `json:"data"`
		Error *string       `json:"error"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&rr); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if rr.Error != nil {
		return nil, errors.New(*rr.Error)
	}

	return rr.Data, nil
}

// Subscribe ...
func (c *Client) Subscribe(subscriptionID string, providerID string) (*Transaction, error) {
	url := fmt.Sprintf("%s/subscriptions/%s/subscribe", c.URL, subscriptionID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if c.token != "" {
		req.Header.Add("Authorization", c.token)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	rr := struct {
		Data  *Transaction `json:"data"`
		Error *string      `json:"error"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&rr); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if rr.Error != nil {
		return nil, errors.New(*rr.Error)
	}

	return rr.Data, nil
}

// Unsubscribe ...
func (c *Client) Unsubscribe(subscriptionID string) (*Transaction, error) {
	url := fmt.Sprintf("%s/subscriptions/%s/unsubscribe", c.URL, subscriptionID)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if c.token != "" {
		req.Header.Add("Authorization", c.token)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	rr := struct {
		Data  *Transaction `json:"data"`
		Error *string      `json:"error"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&rr); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if rr.Error != nil {
		return nil, errors.New(*rr.Error)
	}

	return rr.Data, nil
}

// Refresh ...
func (c *Client) Refresh(subscriptionID string, providerID string) (*Transaction, error) {
	r := struct {
		Refresh *Transaction `json:"refresh"`
	}{}

	r.Refresh = &Transaction{
		ProviderID: providerID,
	}

	b, err := json.Marshal(&r)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/subscriptions/%s/refresh", c.URL, subscriptionID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if c.token != "" {
		req.Header.Add("Authorization", c.token)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	rr := struct {
		Data  *Transaction `json:"data"`
		Error *string      `json:"error"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&rr); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if rr.Error != nil {
		return nil, errors.New(*rr.Error)
	}

	return rr.Data, nil
}

// Verify ...
func (c *Client) Verify(subscriptionID string, transactionID string) (*Transaction, error) {
	r := struct {
		Verify *Transaction `json:"verify"`
	}{}

	r.Verify = &Transaction{
		ID: transactionID,
	}

	b, err := json.Marshal(&r)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/subscriptions/%s/refresh", c.URL, subscriptionID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	if c.token != "" {
		req.Header.Add("Authorization", c.token)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	rr := struct {
		Data  *Transaction `json:"data"`
		Error *string      `json:"error"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&rr); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if rr.Error != nil {
		return nil, errors.New(*rr.Error)
	}

	return rr.Data, nil
}

// ListTransactions ...
func (c *Client) ListTransactions() ([]*Transaction, error) {
	url := fmt.Sprintf("%s/subscriptions/transactions", c.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	rr := struct {
		Data  []*Transaction `json:"data"`
		Error *string        `json:"error"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&rr); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if rr.Error != nil {
		return nil, errors.New(*rr.Error)
	}

	return rr.Data, nil
}

// LastTransaction ...
func (c *Client) LastTransaction() (*Transaction, error) {
	url := fmt.Sprintf("%s/subscriptions/providers", c.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	rr := struct {
		Data  *Transaction `json:"data"`
		Error *string      `json:"error"`
	}{}

	if err := json.NewDecoder(res.Body).Decode(&rr); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if rr.Error != nil {
		return nil, errors.New(*rr.Error)
	}

	return rr.Data, nil
}
