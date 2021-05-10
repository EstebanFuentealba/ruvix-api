package util

import (
	"github.com/cagodoy/ruvix-api/pkg/auth"
	"github.com/cagodoy/ruvix-api/pkg/users"
	uuid "github.com/satori/go.uuid"
)

// FactoryNewAuth ...
func FactoryNewAuth(as auth.Service) (*auth.Response, error) {
	randomUUID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	// After: signup user
	u := &users.User{
		Email:    "fake_email_" + randomUUID.String(),
		Password: "fake_password",
		Name:     "fake_name",
	}

	_, err = as.Signup(u)
	if err != nil {
		return nil, err
	}

	user, err := as.Login(u.Email, u.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
