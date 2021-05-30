package handlers

import (
	"github.com/yukitsune/chameleon/internal/log"
	"gorm.io/gorm"
)

type ValidateHandler struct {
	// Todo: Dependencies go here
	db     *gorm.DB
	logger log.ChameleonLogger
}

func NewValidateHandler(db *gorm.DB, logger log.ChameleonLogger) *ValidateHandler {
	return &ValidateHandler{
		db:     db,
		logger: logger,
	}
}

func (handler *ValidateHandler) Handle(r interface{}) error {
	return nil
}
