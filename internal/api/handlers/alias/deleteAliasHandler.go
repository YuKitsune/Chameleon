package alias

import "github.com/yukitsune/chameleon/internal/api/model"

type DeleteAliasHandler struct {

}

func NewDeleteAliasHandler() *DeleteAliasHandler {
	return &DeleteAliasHandler{}
}

func (handler *DeleteAliasHandler) Handle(req *model.DeleteAliasRequest) error {
	panic("implement me")
}