package template

import (
	"fmt"
	"log"

	"github.com/cagodoy/ruvix-api/pkg/pigeon/email"
	"github.com/cagodoy/ruvix-api/pkg/users"
)

// PasswordChangedTemplate ...
func PasswordChangedTemplate(es *email.Service) func(u *users.User) error {
	return func(u *users.User) error {
		// define template intepolation values
		fpt := PasswordChangedValues{
			Name:    u.Name,
			Company: "MicroAPIs",
		}

		// generate template with interpolation
		str := PasswordChanged(fpt)

		// send email with url
		id, err := es.Send(&email.Message{
			From:     "no-reply@microapis.dev",
			FromName: fpt.Company,
			To:       []string{u.Email},
			Subject:  fmt.Sprintf("[%s]: Recent changes to your Steam account", fpt.Company),
			Text:     str,
			Provider: "sendgrid",
		}, 0)
		if err != nil {
			return err
		}

		log.Printf("Send email for Password changed, email=%s id=%s", u.Email, id)
		return nil
	}
}
