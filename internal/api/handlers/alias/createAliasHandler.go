package alias

import (
	"github.com/yukitsune/chameleon/internal/api/handlers/errors"
	"github.com/yukitsune/chameleon/internal/api/model"
	"gorm.io/gorm"
	"regexp"
)

type CreateAliasHandler struct {
	db *gorm.DB
}

func NewCreateAliasHandler(db *gorm.DB) *CreateAliasHandler {
	return &CreateAliasHandler{db}
}

func (handler *CreateAliasHandler) Handle(req *model.CreateAliasRequest) (*model.Alias, error) {
	alias := req.Alias

	// Ensure the request is valid

	// User must exist
	// Todo: and be the user submitting the request
	var userCount int64
	handler.db.Where(&model.User{}, alias.UserID).Count(&userCount)
	if userCount == 0 {
		return nil, errors.NewEntityInvalidError(&alias.Model, "user does not exist")
	}

	// Recipient must not be empty
	if len(alias.Username) == 0 {
		return nil, errors.NewEntityInvalidError(&alias,"username must not be empty")
	}

	// Sender whitelist pattern must be valid regex
	_, err := regexp.Compile(alias.SenderWhitelistPattern)
	if err != nil {
		return nil, errors.NewEntityInvalidErrorFromErr(&alias, err)
	}

	// Ensure no duplicate entries exist
	var dupe *model.Alias
	handler.db.Where(&model.Alias{
		UserID: alias.UserID,
		Username: alias.Username,
		SenderWhitelistPattern: alias.SenderWhitelistPattern,
	}).First(&dupe)
	if dupe != nil {
		return nil, errors.NewEntityExistsError(&alias)
	}

	// All good, create the record
	result := handler.db.Create(&alias)
	if result.Error != nil {
		return nil, result.Error
	}

	return &alias, nil
}