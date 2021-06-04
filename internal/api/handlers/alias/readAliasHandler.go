package alias

import "github.com/yukitsune/chameleon/internal/api/model"

type ReadAliasHandler struct {

}

func NewReadAliasHandler() *ReadAliasHandler {
	return &ReadAliasHandler{}
}

func (handler *ReadAliasHandler) Handle(req *model.GetAliasRequest) (*model.Alias, error) {
	panic("implement me")
}