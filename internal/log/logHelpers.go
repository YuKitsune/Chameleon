package log

import (
	"github.com/sirupsen/logrus"
)

func (log *logrus.Logger) IsDebug() bool {
	return log.Level == logrus.DebugLevel || log.Level == logrus.TraceLevel
}
