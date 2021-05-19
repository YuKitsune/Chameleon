package api

import "fmt"

type DbConfig struct {
	Host     string    `yaml:"host"`
	Port int `yaml:"port"`
	User  string `yaml:"user"`
	Password  string `yaml:"password"`
	Database  string `yaml:"database"`
}

func (c *DbConfig) ConnectionString() string {
	// Todo: Hand this to a provider so we can use different databases
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=verify-full", c.Host, c.Port, c.User, c.Password, c.Database)
}

func (c *DbConfig) SetDefaults() error {

	// Todo: Hand this to a provider so we can use different databases

	return nil
}
