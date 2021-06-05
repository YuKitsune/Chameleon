package alias

import "github.com/yukitsune/chameleon/internal/api/model"

type DeleteAliasHandler struct {

}

func NewDeleteAliasHandler() *DeleteAliasHandler {
	return &DeleteAliasHandler{}
}

func (handler *DeleteAliasHandler) Handle(req *model.DeleteAliasRequest) (bool, error) {
	panic("implement me")
}