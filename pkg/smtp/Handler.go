package smtp

import (
	"github.com/yukitsune/chameleon/internal/log"
)

type Handler interface {
	ValidateRcpt(e *Envelope, logger log.ChameleonLogger) error
	Handle(e *Envelope, logger log.ChameleonLogger) Result
}
