package mandrill

import (
	"log"

	"github.com/cagodoy/ruvix-api/pkg/pigeon/email"
	"github.com/keighl/mandrill"
)

const (
	Name = "mandrill"
)

// Mandrill ...
type Mandrill struct {
	key string
}

// New ...
func New(key string) *Mandrill {
	return &Mandrill{
		key: key,
	}
}

// Approve ...
func (p *Mandrill) Approve(*email.Message) error {
	return nil
}

// Deliver ...
func (p *Mandrill) Deliver(m *email.Message) error {
	// create client
	client := mandrill.ClientWithKey(p.key)

	// if not has HTML, set text message to HTML
	if m.HTML == "" {
		m.HTML = m.Text
	}

	// prepare message
	email := &mandrill.Message{
		FromEmail: m.From,
		FromName:  m.FromName,
		Subject:   m.Subject,
		HTML:      m.HTML,
		Text:      m.Text,
	}

	// add email recipient
	email.AddRecipient(m.To[0], m.To[0], "to")

	// send email
	response, err := client.MessagesSend(email)
	if err != nil {
		return err
	}

	log.Println(response[0].Email)
	log.Println(response[0].Id)
	log.Println(response[0].Status)
	log.Println(response[0].RejectionReason)

	return nil
}
