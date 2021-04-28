package client

import "github.com/yukitsune/chameleon/pkg/smtp"

type ChameleonClient interface {
	Validate(sender string, recipient string) error
	Handle(e *smtp.Envelope) error
}
