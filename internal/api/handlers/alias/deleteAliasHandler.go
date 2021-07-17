package alias

import (
	"github.com/yukitsune/chameleon/internal/api/model"
	"github.com/yukitsune/chameleon/internal/log"
	"gorm.io/gorm"
)

type DeleteAliasHandler struct {
	db *gorm.DB
	log log.ChameleonLogger
}

func NewDeleteAliasHandler(db *gorm.DB, log log.ChameleonLogger) *DeleteAliasHandler {
	return &DeleteAliasHandler{db, log}
}

func (handler *DeleteAliasHandler) Handle(req *model.DeleteAliasRequest) (bool, error) {
	res := handler.db.Delete(&req.Alias)
	if res.Error != nil {
		return false, res.Error
	}

	return true, nil
}
