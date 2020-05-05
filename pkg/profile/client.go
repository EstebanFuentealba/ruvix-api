package profile

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	uluru "github.com/jmlopezz/uluru-api"
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
func NewClient(address string, opts uluru.ClientOptions) (*Client, error) {
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

type profileRespose struct {
	Data  *Profile `json:"data"`
	Error *string  `json:"error"`
}

type profileRequest struct {
	Profile *Profile `json:"profile"`
}

// Get ...
func (c *Client) Get() (*Profile, error) {
	url := fmt.Sprintf("%s/profile", c.environment)

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

	p := &profileRespose{}
	if err := json.NewDecoder(res.Body).Decode(p); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if p.Error != nil {
		return nil, errors.New(*p.Error)
	}

	return p.Data, nil
}

// Update ...
func (c *Client) Update(p *Profile) (*Profile, error) {
	r := &profileRequest{
		Profile: p,
	}

	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/profile", c.environment)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(b)))
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

	pp := &profileRespose{}
	if err := json.NewDecoder(res.Body).Decode(pp); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if pp.Error != nil {
		return nil, errors.New(*pp.Error)
	}

	return pp.Data, nil
}

// Create ...
func (c *Client) Create(p *Profile) (*Profile, error) {
	r := &profileRequest{
		Profile: p,
	}

	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/profile", c.environment)
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

	pp := &profileRespose{}
	if err := json.NewDecoder(res.Body).Decode(pp); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if pp.Error != nil {
		return nil, errors.New(*pp.Error)
	}

	return pp.Data, nil
}
