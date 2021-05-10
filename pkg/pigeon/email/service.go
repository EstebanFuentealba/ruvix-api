package email

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/cagodoy/ruvix-api/pkg/pigeon/message"
	"github.com/cagodoy/ruvix-api/pkg/pigeon/scheduler"
)

type Provider interface {
	Approve(*Message) error
	Deliver(*Message) error
}

// Service ...
type Service struct {
	Scheduler scheduler.Service
}

// NewService ...
func NewService(pq *scheduler.PriorityQueue, provider Provider) (*Service, error) {
	// Init DB
	boltDst, err := message.NewBoltDatastore("messages.db")
	if err != nil {
		return nil, err
	}

	// initialize message store
	ms, err := message.NewMessageStore(boltDst)
	if err != nil {
		return nil, err
	}

	return &Service{
		Scheduler: scheduler.New(pq, ms, approve(provider), deliver(provider)),
	}, nil
}

func (s *Service) Send(e *Message, delay int64) (string, error) {
	// validates email existence
	if e == nil {
		return "", errors.New("invalid email")
	}

	// validates email params
	// TODO(ca): check how verify e.Text, e.HTML, e.ReplyTo and e.Status
	if e.From == "" || e.FromName == "" || e.To[0] == "" || e.Subject == "" || e.Provider == "" {
		return "", errors.New("invalid email params")
	}

	// validates delay param
	if delay < 0 {
		return "", errors.New("invalid delay")
	}

	// pase email struct
	b, err := json.Marshal(e)
	if err != nil {
		return "", err
	}

	e.Status = message.Pending
	id, err := s.Scheduler.Put(delay, "email", "sendgrid", string(b), e.Status)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func approve(provider Provider) func(content string) (bool, error) {
	return func(content string) (bool, error) {
		log.Println(fmt.Sprintf("[Email][Approve][Request] content = %v", content))

		if content == "" {
			err := errors.New("invalid message content")
			log.Println(fmt.Sprintf("[Email][Approve][Error] error = %v", err))
			return false, err
		}

		m := new(Message)
		err := json.Unmarshal([]byte(content), m)
		if err != nil {
			log.Println(fmt.Sprintf("[Email][Approve][Error] error = %v", err))
			return false, err
		}

		err = provider.Approve(m)
		if err != nil {
			return false, err
		}

		log.Println(fmt.Sprintf("[Email][Approve][Response] message = %v", m))

		return true, nil
	}
}

func deliver(provider Provider) func(content string) error {
	return func(content string) error {
		log.Println(fmt.Sprintf("[Email][Deliver][Request] content = %v", content))

		if content == "" {
			err := errors.New("invalid message content")
			log.Println(fmt.Sprintf("[Email][Deliver][Error] error = %v", err))
			return err
		}

		m := new(Message)
		err := json.Unmarshal([]byte(content), m)
		if err != nil {
			log.Println(fmt.Sprintf("[Email][Deliver][Error] error = %v", err))
			return err
		}

		err = provider.Approve(m)
		if err != nil {
			return err
		}

		log.Println(fmt.Sprintf("[Email][Approve][Response] message = %v", m))

		return nil
	}
}
