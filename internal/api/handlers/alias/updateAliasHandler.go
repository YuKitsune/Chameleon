package alias

import "github.com/yukitsune/chameleon/internal/api/model"

type UpdateAliasHandler struct {

}

func NewUpdateAliasHandler() *UpdateAliasHandler {
	return &UpdateAliasHandler{}
}

func (handler *UpdateAliasHandler) Handle(req *model.UpdateAliasRequest) (*model.Alias, error) {
	panic("implement me")
}