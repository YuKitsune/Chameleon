package api

import "fmt"

type DbConfig struct {
	Host     string    `mapstructure:"host"`
	Port int `mapstructure:"port"`
	User  string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
}

func (c *DbConfig) ConnectionString() string {
	// Todo: Hand this to a provider so we can use different databases
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=verify-full", c.Host, c.Port, c.User, c.Password, c.Database)
}

func (c *DbConfig) SetDefaults() error {

	// Todo: Hand this to a provider so we can use different databases

	return nil
}
