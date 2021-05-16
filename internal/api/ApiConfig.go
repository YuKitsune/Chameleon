package api

import "fmt"

type ApiConfig struct {
	Port     int    `yaml:"port"`
	CertFile string `yaml:"certificate-file"`
	KeyFile  string `yaml:"key-file"`
}

func (c ApiConfig) GetAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}

func (c *ApiConfig) SetDefaults() error {
	if c.Port == 0 {
		c.Port = 80
	}

	return nil
}
