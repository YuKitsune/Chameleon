package log

import "github.com/sirupsen/logrus"

func NewLogrusLogger(cfg *Config) (*logrus.Logger, error) {
	logger := logrus.New()

	lvl, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		return nil, err
	}

	logger.SetLevel(lvl)

	switch cfg.Formatter {
	case TextFormatter:
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:               true,
			PadLevelText:              true,
		})
	case JsonFormatter:
		logger.SetFormatter(&logrus.JSONFormatter{})
	}

	return logger, nil
}
