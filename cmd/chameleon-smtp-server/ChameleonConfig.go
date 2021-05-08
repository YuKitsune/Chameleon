package main

import (
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/pkg/smtp"
)

type ChameleonSmtpServerConfig struct {
	ApiUrl  string             `yaml:"chameleon-api-base-url"`
	Smtp    *smtp.ServerConfig `yaml:"smtp"`
	Logging *log.LogConfig     `yaml:"log"`
}

func (c *ChameleonSmtpServerConfig) SetDefaults() error {
	err := c.Smtp.SetDefaults()
	if err != nil {
		return err
	}

	err = c.Logging.SetDefaults()
	if err != nil {
		return err
	}

	return nil
}