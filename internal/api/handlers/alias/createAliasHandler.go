package alias

import (
	"github.com/yukitsune/chameleon/internal/api/handlers/errors"
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/log"
	"gorm.io/gorm"
	"regexp"
)

type CreateAliasHandler struct {
	db *gorm.DB
	log log.ChameleonLogger
}

func NewCreateAliasHandler(db *gorm.DB, log log.ChameleonLogger) *CreateAliasHandler {
	return &CreateAliasHandler{db, log}
}

func (handler *CreateAliasHandler) Handle(req *model.CreateAliasRequest) (*model.Alias, error) {
	alias := req.Alias

	// Ensure the request is valid

	// Todo: User must exist and be the user submitting the request
	//var userCount int64
	//handler.db.Where(&model.User{}, alias.UserID).Count(&userCount)
	//if userCount == 0 {
	//	return nil, errors.NewEntityInvalidError(&alias.Model, "user does not exist")
	//}

	// Recipient must not be empty
	if len(alias.Username) == 0 {
		return nil, errors.NewEntityInvalidError(&alias, "username must not be empty")
	}

	// Sender whitelist pattern must be valid regex
	_, err := regexp.Compile(alias.SenderWhitelistPattern)
	if err != nil {
		return nil, errors.NewEntityInvalidErrorFromErr(&alias, err)
	}

	// Ensure no duplicate entries exist
	var dupes int64
	handler.db.Where(&model.Alias{
		Username:               alias.Username,
		SenderWhitelistPattern: alias.SenderWhitelistPattern,
	}).Count(&dupes)
	if dupes != 0 {
		return nil, errors.NewEntityExistsError(&alias)
	}

	// All good, create the record
	result := handler.db.Create(&alias)
	if result.Error != nil {
		return nil, result.Error
	}

	return &alias, nil
}
