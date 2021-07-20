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

func (handler *DeleteAliasHandler) Handle(req *model.DeleteAliasRequest) (*model.DeleteAliasResponse, error) {
	res := handler.db.Delete(&model.Alias{}, req.Alias.ID)
	if res.Error != nil {
		return &model.DeleteAliasResponse{Deleted: false}, res.Error
	}

	return &model.DeleteAliasResponse{Deleted: true}, nil
}
