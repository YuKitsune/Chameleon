package log

type LogConfig struct {
	Directory string `json:"directory"`
	Level     string `json:"level"`
}

func (l *LogConfig) SetDefaults() error {
	if l.Directory == "" {
		l.Directory = "./log"
	}

	if l.Level == "" {
		l.Level = InfoLevel.String()
	}

	return nil
}
