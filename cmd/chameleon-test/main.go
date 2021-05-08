package main

import (
	"fmt"
	"github.com/yukitsune/chameleon/cmd"
	"net/smtp"
)

func main() {

	logger := cmd.MakeLogger("trace", "./logs")

	fmt.Println("Hello, World!")

	host := "127.0.0.1:2525"
	client, err := smtp.Dial(host)
	if err != nil {
		return
	}

	err = client.Rcpt("test@relay.chameleon.io")
	if err != nil {
		return
	}

	// Todo: this part won't work, need to send mail using a different mail server
	// auth := smtp.PlainAuth("test_id", "test_username", "test_password", host)

	to := []string{"recipient@relay.chameleon.io"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	err = smtp.SendMail(host, nil, "sender@example.org", to, msg)
	if err != nil {
		logger.Fatal(err)
	}

	cmd.WaitForShutdownSignal(logger)
}
