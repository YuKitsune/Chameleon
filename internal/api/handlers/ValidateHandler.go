package handlers

import (
	"github.com/yukitsune/chameleon/internal/log"
	"net/http"
)

type ValidateHandler struct {
	// Todo: Dependencies go here
	logger log.ChameleonLogger
}

func NewValidateHandler(logger log.ChameleonLogger) *ValidateHandler {
	return &ValidateHandler{logger: logger}
}

func (handler *ValidateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	handler.logger.Info("Validating recipient")
}