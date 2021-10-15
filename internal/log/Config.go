package log

import "github.com/sirupsen/logrus"

const (
	TextFormatter int = iota
	JsonFormatter
)

type Config struct {
	Level     string `mapstructure:"level"`
	Formatter int `mapstructure:"formatter"`
}

func (l *Config) SetDefaults() error {
	if l.Level == "" {
		l.Level = logrus.InfoLevel.String()
	}

	return nil
}
