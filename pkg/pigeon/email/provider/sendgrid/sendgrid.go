package sendgrid

import (
	"errors"
	"log"

	"github.com/cagodoy/ruvix-api/pkg/pigeon/email"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	Name = "sendgrid"
)

// Sendgrid ...
type Sendgrid struct {
	key string
}

// New ...
func New(key string) *Sendgrid {
	return &Sendgrid{
		key: key,
	}
}

// Approve ...
func (p *Sendgrid) Approve(*email.Message) error {
	return nil
}

// Deliver ...
func (p *Sendgrid) Deliver(m *email.Message) error {
	// define from and to values
	from := mail.NewEmail(m.FromName, m.From)
	to := mail.NewEmail(m.To[0], m.To[0])

	// if not has HTML, set text message to HTML
	if m.HTML == "" {
		m.HTML = m.Text
	}

	// prepare single email
	message := mail.NewSingleEmail(from, m.Subject, to, m.Text, m.HTML)

	// create client
	client := sendgrid.NewSendClient(p.key)

	// send message
	response, err := client.Send(message)
	if err != nil {
		return err
	}

	log.Println("=================")
	log.Println("=================")
	log.Println(response.StatusCode)
	log.Println(response.Body)
	log.Println(response.Headers)
	log.Println("=================")
	log.Println("=================")

	if response.StatusCode >= 400 {
		return errors.New(response.Body)
	}

	return nil
}
