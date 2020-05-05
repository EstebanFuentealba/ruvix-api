package util

import (
	auth "github.com/microapis/authentication-api"
	authclient "github.com/microapis/authentication-api/client"
	"github.com/microapis/users-api"
	uuid "github.com/satori/go.uuid"
)

// FactoryNewAuth ...
func FactoryNewAuth(authAddr string) (*auth.Response, error) {
	ac, err := authclient.New(authAddr)
	if err != nil {
		return nil, err
	}

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

	_, err = ac.Signup(u)
	if err != nil {
		return nil, err
	}

	user, err := ac.Login(u.Email, u.Password)
	if err != nil {
		return nil, err
	}

	return user, nil
}
