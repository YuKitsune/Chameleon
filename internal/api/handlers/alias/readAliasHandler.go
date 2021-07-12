package alias

import (
	"github.com/yukitsune/chameleon/internal/api/handlers/errors"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/log"
	"gorm.io/gorm"
	"regexp"
)

type ReadAliasHandler struct {
	db *gorm.DB
	log log.ChameleonLogger
}

func NewReadAliasHandler(db *gorm.DB, log log.ChameleonLogger) *ReadAliasHandler {
	return &ReadAliasHandler{db, log}
}

func (handler *ReadAliasHandler) Handle(req *model.GetAliasRequest) (*model.Alias, error) {

	var allAliasesForRecipient []model.Alias
	handler.db.Where(&model.Alias{
		Username: req.Recipient,
	}).Find(&allAliasesForRecipient)

	for _, alias := range allAliasesForRecipient {
		isMatch, err := regexp.MatchString(alias.SenderWhitelistPattern, req.Sender)
		if err != nil {
			return nil, err
		}

		if isMatch {
			return &alias, nil
		}
	}

	return nil, errors.NewEntityNotFoundError(req)
}
