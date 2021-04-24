package smtp

import (
	"github.com/sirupsen/logrus"
)

type Handler interface {
	ValidateRcpt(e *Envelope, logger *logrus.Logger) error
	Handle(e *Envelope, logger *logrus.Logger) Result
}
