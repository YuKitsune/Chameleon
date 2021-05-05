package log

const (
	LogDirKey = "log-dir"
	LogLevelKey = "log-level"
)

type LogConfig struct {
	LogDirectory string `json:"log-dir"`
	Level string `json:"log-level"`
}