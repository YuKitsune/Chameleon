package smtp

import "log"

type Handler interface {
	validateRctp() bool
	handle(logger log.Logger) bool
}