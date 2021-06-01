package mediatorHandlers

import (
	"github.com/yukitsune/chameleon/internal/log"
)

type MailHandler struct {
	// Todo: Dependencies go here
	logger log.ChameleonLogger
}

func NewMailHandler(logger log.ChameleonLogger) *MailHandler {
	return &MailHandler{logger: logger}
}

func (handler *MailHandler) Handle(r interface{}) error {
	handler.logger.Info("Handling mail")
	return nil
}
