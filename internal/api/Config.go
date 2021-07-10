package api

import (
	"fmt"
)

type Config struct {
	Port     int       `mapstructure:"port"`
	CertFile string    `mapstructure:"certificate-file"`
	KeyFile  string    `mapstructure:"key-file"`
	Database *DbConfig `mapstructure:"database"`
}

func (c Config) GetAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}

func (c *Config) SetDefaults() error {
	if c.Port == 0 {
		c.Port = 80
	}

	if c.Database == nil {
		dbConfig := &DbConfig{}
		err := dbConfig.SetDefaults()
		if err != nil {
			return err
		}

		c.Database = dbConfig
	}

	return nil
}
