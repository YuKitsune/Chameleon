package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
)

type LogrusChameleonLogger struct {
	logrus.Logger
}

func (l *LogrusChameleonLogger) GetLevel() Level {
	logrusLevel := l.Logger.GetLevel()
	levelInt := uint8(logrusLevel)
	chameleonLevel := Level(levelInt)
	return chameleonLevel
}

func (l *LogrusChameleonLogger) WithField(key string, value interface{}) ChameleonLogger {
	return wrapEntry(l.Logger.WithField(key, value))
}

func (l *LogrusChameleonLogger) WithFields(fields Fields) ChameleonLogger {
	return wrapEntry(l.Logger.WithFields(logrus.Fields(fields)))
}

func (l *LogrusChameleonLogger) WithError(err error) ChameleonLogger {
	return wrapEntry(l.Logger.WithError(err))
}

func (l *LogrusChameleonLogger) IsDebug() bool {
	return l.GetLevel() == DebugLevel
}

func wrapEntry(entry *logrus.Entry) *LogrusChameleonLogger {
	return &LogrusChameleonLogger{
		*entry.Logger,
	}
}

func NewLogrusChameleonLogger(filename string, level Level) (*LogrusChameleonLogger, error) {

	writer, err := getLogWriter(filename)
	if err != nil {
		return nil, err
	}

	logrusLevel, err := logrus.ParseLevel(level.String())
	if err != nil {
		return nil, err
	}

	logger := &LogrusChameleonLogger{
		logrus.Logger{
			Out:       writer,
			Formatter: new(logrus.TextFormatter),
			Hooks:     make(logrus.LevelHooks),
			Level:     logrusLevel,
		},
	}

	return logger, nil
}

func getLogWriter(filename string) (io.Writer, error) {

	// Ensure the directory exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		err := os.MkdirAll(filepath.Dir(filename), 0700)
		if err != nil {
			return nil, err
		}
	}

	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	writer := io.MultiWriter(f, os.Stdout)

	return writer, nil
}

func LogrusLoggerProvider(file string, level Level) (ChameleonLogger, error) {
	logger, err := NewLogrusChameleonLogger(file, level)
	if err != nil {
		return nil, err
	}

	return logger, nil
}
