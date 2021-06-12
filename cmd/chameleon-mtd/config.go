package main

import (
	"github.com/yukitsune/chameleon/internal/log"
	"github.com/yukitsune/chameleon/pkg/smtp"
)

type ChameleonMtdConfig struct {
	ApiUrl  string       `mapstructure:"api-url"`
	Smtp    *smtp.Config `mapstructure:"smtp"`
	Logging *log.Config  `mapstructure:"log"`
}

func (c *ChameleonMtdConfig) SetDefaults() error {
	if c.Smtp == nil {
		c.Smtp = &smtp.Config{}
	}
	err := c.Smtp.SetDefaults()
	if err != nil {
		return err
	}

	if c.Logging == nil {
		c.Logging = &log.Config{}
	}
	err = c.Logging.SetDefaults()
	if err != nil {
		return err
	}

	return nil
}