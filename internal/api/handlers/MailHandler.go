package handlers

import (
	"github.com/yukitsune/chameleon/internal/log"
	"net/http"
)

type MailHandler struct {
	// Todo: Dependencies go here
	logger log.ChameleonLogger
}

func NewMailHandler(logger log.ChameleonLogger) *MailHandler {
	return &MailHandler{logger: logger}
}

func (handler *MailHandler) Handle(w http.ResponseWriter, r *http.Request) {
	handler.logger.Info("Handling mail")
}