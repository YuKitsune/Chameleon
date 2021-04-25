package log

import (
	"fmt"
	"path/filepath"
	"time"
)

type fileNameProviderFunc func(time time.Time) string
type loggerProviderFunc func(file string, level Level) (ChameleonLogger, error)

type LogFactory struct {
	directory    string
	level        Level
	makeFileName fileNameProviderFunc
	makeLogger   loggerProviderFunc
}

func NewLogFactory(dir string, level Level, fileNameProvider fileNameProviderFunc, loggerProvider loggerProviderFunc) *LogFactory {
	path := filepath.Clean(dir)
	return &LogFactory{
		directory:    path,
		level:        level,
		makeFileName: fileNameProvider,
		makeLogger:   loggerProvider,
	}
}

func (f *LogFactory) Make() (ChameleonLogger, error) {
	fileName := f.makeFileName(time.Now())
	fullFileName := fmt.Sprintf("%s/%s", f.directory, fileName)

	logger, err := f.makeLogger(fullFileName, f.level)
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func DefaultFileNameProvider(time time.Time) string {
	return fmt.Sprintf("Chameleon_Log_%s.log", time.Format("2006-02-01"))
}
