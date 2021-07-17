package alias

import (
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/log"
	"gorm.io/gorm"
)

type UpdateAliasHandler struct {
	db *gorm.DB
	log log.ChameleonLogger
}

func NewUpdateAliasHandler(db *gorm.DB, log log.ChameleonLogger) *UpdateAliasHandler {
	return &UpdateAliasHandler{db, log}
}

func (handler *UpdateAliasHandler) Handle(req *model.UpdateAliasRequest) (*model.Alias, error) {

	var alias model.Alias
	res := handler.db.First(&alias, req.Alias.ID)
	if res.Error != nil {
		return nil, res.Error
	}

	res = handler.db.Save(alias)
	if res.Error != nil {
		return nil, res.Error
	}

	return &req.Alias, nil
}
