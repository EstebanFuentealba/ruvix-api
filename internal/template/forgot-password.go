package template

import (
	"fmt"
	"log"

	"github.com/cagodoy/ruvix-api/pkg/pigeon/email"
	"github.com/cagodoy/ruvix-api/pkg/users"
)

// ForgotPasswordTemplate ...
func ForgotPasswordTemplate(es *email.Service) func(u *users.User, token string) error {
	return func(u *users.User, token string) error {
		// define template intepolation values
		fpt := ForgotPasswordValues{
			Email:      u.Email,
			TokenURL:   fmt.Sprintf("https://www.microapis.dev/password-reset?token=%s", token),
			ExpireTime: "5 minutes",
			Company:    "MicroAPIs",
		}

		// generate template with interpolation
		str := ForgotPassword(fpt)

		// send email with token and url
		id, err := es.Send(&email.Message{
			From:     "no-reply@microapis.dev",
			FromName: u.Name,
			To:       []string{u.Email},
			Subject:  fmt.Sprintf("[%s]: Instructions for changing your password", fpt.Company),
			Text:     str,
			Provider: "sendgrid",
		}, 0)
		if err != nil {
			return err
		}

		log.Printf("Send email for Forgot password, email=%s id=%s", u.Email, id)
		return nil
	}
}
