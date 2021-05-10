package ses

import (
	"html"

	"github.com/cagodoy/ruvix-api/pkg/pigeon/email"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	Name = "ses"
)

// SES ...
type SES struct {
	region    string
	key       string
	secretKey string
}

// New ...
func New(region string, key string, secretKey string) *SES {
	return &SES{
		region:    region,
		key:       key,
		secretKey: secretKey,
	}
}

// Approve ...
func (p *SES) Approve(*email.Message) error {
	return nil
}

// Deliver ...
func (p *SES) Deliver(m *email.Message) error {
	// define aws config credentials
	config := &aws.Config{
		Region:      aws.String(p.region),
		Credentials: credentials.NewStaticCredentials(p.key, p.secretKey, ""),
	}

	// define aws session
	session := session.New(config)

	// define ses instance
	sesClient := ses.New(session)

	// if not has HTML, set text message to HTML
	if m.HTML == "" {
		m.HTML = m.Text
	}

	// unescape html source
	m.HTML = html.UnescapeString(m.HTML)

	// prepare message with email values
	msg := &ses.Message{
		Subject: &ses.Content{
			Charset: aws.String("utf-8"),
			Data:    &m.Subject,
		},
		Body: &ses.Body{
			Html: &ses.Content{
				Charset: aws.String("utf-8"),
				Data:    &m.HTML,
			},
			Text: &ses.Content{
				Charset: aws.String("utf-8"),
				Data:    &m.Text,
			},
		},
	}

	// define emails destinations
	dest := &ses.Destination{
		ToAddresses: aws.StringSlice(m.To),
	}

	// send emails to destinations
	_, err := sesClient.SendEmail(&ses.SendEmailInput{
		Source:           &m.From,
		Destination:      dest,
		Message:          msg,
		ReplyToAddresses: aws.StringSlice(m.ReplyTo),
	})
	if err != nil {
		return err
	}

	return nil
}
