package main

import "github.com/yukitsune/chameleon/pkg/smtp"

type ChameleonSmtpServerConfig struct {
	ApiUrl string `json:"chameleon-api-base-url"`
	AllowedHosts []string `json:"allowed-hosts"` // Todo: this seems like something for SMTP config
	SmtpConfig smtp.ServerConfig `json:"smtp"`
}
