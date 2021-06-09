package users

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func Post(url string, u interface{}) (*User, error) {
	postBody, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &struct {
		Data  *User       `json:"data,omitempty"`
		Meta  interface{} `json:"meta,omitempty"`
		Error string      `json:"error,omitempty"`
	}{}

	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}

	return r.Data, nil
}

func Get(url string) (*User, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := &struct {
		Data  *User       `json:"data,omitempty"`
		Meta  interface{} `json:"meta,omitempty"`
		Error string      `json:"error,omitempty"`
	}{}

	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	if r.Error != "" {
		return nil, errors.New(r.Error)
	}

	return r.Data, nil
}
