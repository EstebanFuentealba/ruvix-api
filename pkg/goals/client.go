package goals

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
	client      http.Client
	environment string
	token       string
}

// NewClient ...
func NewClient(address string, opts urulu.ClientOptions) (*Client, error) {
	// TODO(ca): check on authorization if token is valid
	t := opts.Token

	c := http.Client{
		Timeout: time.Duration(50 * time.Second),
	}

	var env string
	if opts.Environment == "production" {
		env = baseURLProduction
	} else if opts.Environment == "staging" {
		env = baseURLStaging
	} else if opts.Environment == "development" {
		env = baseURLDevelopment
	} else {
		env = baseURLDevelopment
	}

	return &Client{
		client:      c,
		environment: env,
		token:       t,
	}, nil
}

// List ...
func (c *Client) List() ([]*Goal, error) {
	url := fmt.Sprintf("%s/goals", c.environment)

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
		Data  []*Goal `json:"data"`
		Error *string `json:"error"`
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

// Create ...
func (c *Client) Create(g *Goal) (*Goal, error) {
	r := struct {
		Goal *Goal `json:"goal"`
	}{
		Goal: g,
	}

	b, err := json.Marshal(&r)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/goals", c.environment)
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
		Data  *Goal   `json:"data"`
		Error *string `json:"error"`
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

// GetRetirement ...
func (c *Client) GetRetirement() (*RetirementGoal, error) {
	url := fmt.Sprintf("%s/goals/retirements", c.environment)

	req, err := http.NewRequest("GET", url, nil)
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
		Data  *RetirementGoal `json:"data"`
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

// CreateRetirement ...
func (c *Client) CreateRetirement(ret *RetirementGoal) (*RetirementGoal, error) {
	r := struct {
		Retirement *RetirementGoal `json:"retirement"`
	}{
		Retirement: ret,
	}

	b, err := json.Marshal(&r)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/goals/retirements", c.environment)
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
		Data  *RetirementGoal `json:"data"`
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
