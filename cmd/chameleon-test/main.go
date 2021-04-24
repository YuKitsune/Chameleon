package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {

	fmt.Println("Hello, World!")

	host := "127.0.0.1:2525"
	client, err := smtp.Dial(host)
	if err != nil {
		return
	}

	err = client.Rcpt("test@test.com")
	if err != nil {
		return
	}

	auth := smtp.PlainAuth("test_id", "test_username", "test_password", host)

	to := []string{"recipient@example.net"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")

	err = smtp.SendMail(host, auth, "sender@example.org", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}