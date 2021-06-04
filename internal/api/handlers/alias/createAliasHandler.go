package alias

import "github.com/yukitsune/chameleon/internal/api/model"

type CreateAliasHandler struct {

}

func NewCreateAliasHandler() *CreateAliasHandler {
	return &CreateAliasHandler{}
}

func (handler *CreateAliasHandler) Handle(req *CreateAliasHandler) (*model.Alias, error) {
	panic("implement me")
}