package api

import "fmt"

type ApiConfig struct {
	Port     int    `yaml:"port"`
	CertFile string `yaml:"certificate-file"`
	KeyFile  string `yaml:"key-file"`
	Database *DbConfig `yaml:"database"`
}

func (c ApiConfig) GetAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}

func (c *ApiConfig) SetDefaults() error {
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
