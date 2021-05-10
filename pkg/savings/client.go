package savings

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	ruvixapi "github.com/cagodoy/ruvix-api"
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
func NewClient(address string, opts ruvixapi.ClientOptions) (*Client, error) {
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

type institutionsResponse struct {
	Data  []*Institution `json:"data"`
	Error *string        `json:"error"`
}

// ListInstitutions ...
func (c *Client) ListInstitutions() ([]*Institution, error) {
	url := fmt.Sprintf("%s/savings/institutions", c.environment)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	rr := &institutionsResponse{}
	if err := json.NewDecoder(res.Body).Decode(rr); err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if rr.Error != nil {
		return nil, errors.New(*rr.Error)
	}

	return rr.Data, nil
}

// CreateInstitution ...
func (c *Client) CreateInstitution(i *Institution) (*Institution, error) {
	r := &createInstitutionRequest{
		Institution: i,
	}

	b, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/savings/institutions", c.environment)
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
		Data  *Institution `json:"data"`
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
