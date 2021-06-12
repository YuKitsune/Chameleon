package main

import (
	"github.com/yukitsune/chameleon/internal/api"
	"github.com/yukitsune/chameleon/internal/log"
)

type ChameleonApiConfig struct {
	Api     *api.Config `mapstructure:"api"`
	Logging *log.Config `mapstructure:"log"`
}

func (c *ChameleonApiConfig) SetDefaults() error {
	if c.Api == nil {
		c.Api = &api.Config{}
	}
	err := c.Api.SetDefaults()
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