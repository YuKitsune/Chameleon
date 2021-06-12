package log

type Config struct {
	Directory string `mapstructure:"directory"`
	Level     string `mapstructure:"level"`
}

func (l *Config) SetDefaults() error {
	if l.Directory == "" {
		l.Directory = "./log"
	}

	if l.Level == "" {
		l.Level = InfoLevel.String()
	}

	return nil
}
