package handlers

import (
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/log"
	"gorm.io/gorm"
	"net/http"
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

func (handler *ValidateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	handler.logger.Info("Validating recipient")
	sender, err := GetParameter("sender", r)
	if err != nil {
		// Todo: Write error response
	}

	recipient, err := GetParameter("recipient", r)
	if err != nil {
		// Todo: Write error response
	}

	// Todo: Note: Any errors from here on should indicate that the recipient doesn't exist

	var alias model.Alias
	res := handler.db.Where("username = ?", recipient).First(&alias)
	if res.Error == gorm.ErrRecordNotFound {
		// Todo: Write no recipient response
		return
	}

	senderIsAllowed, err := alias.SenderIsAllowed(*sender)
	if err != nil {
		// Todo: Write no recipient response
		// 	Log the error
		return
	}

	if !*senderIsAllowed {
		// Todo: Write no recipient response
		return
	}

	// Made it this far, we're valid
	// Todo: Write valid response
}
