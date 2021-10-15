package handlers

import (
	"github.com/sirupsen/logrus"
	"github.com/yukitsune/chameleon/pkg/client"
	"github.com/yukitsune/chameleon/pkg/smtp"
	"net/url"
)

type DefaultHandler struct {
	client client.HttpChameleonClient
}

func NewDefaultHandler(baseApiUrl *url.URL, logger *logrus.Logger) *DefaultHandler {
	chameleonClient := client.NewHttpChameleonClient(baseApiUrl, logger)
	return &DefaultHandler{
		client: chameleonClient,
	}
}

func (d DefaultHandler) ValidateRcpt(e *smtp.Envelope, logger *logrus.Logger) error {

	// Todo: Check each recipient and ensure at least one of them is valid
	sender := e.MailFrom.String()
	recipient := e.RcptTo[len(e.RcptTo)-1].String()

	return d.client.Validate(sender, recipient)

	// Todo: Run these checks server-side
	// 	/api/validate?sender=abc@123.com?recipient=alias@chameleon.io
	// What makes a recipient (address alias) valid?
	// 1. The recipient exists
	// 2. The owner of the recipient has not blacklisted the sender
	//    - By exact address
	//	  - By domain
	// 3. The recipient has not blacklisted the sender
	//    - By exact address
	//	  - By domain
	// 4. The recipient has added the sender to a whitelist
	//    - By exact address
	//	  - By domain
}

func (d DefaultHandler) Handle(e *smtp.Envelope, logger *logrus.Logger) smtp.Result {
	err := d.client.Handle(e)
	if err != nil {
		return smtp.NewResult(smtp.Canned.FailBackendTransaction)
	}

	// Todo: Not sure if this is the right result to return
	return smtp.NewResult(smtp.Canned.SuccessMessageQueued)
}
