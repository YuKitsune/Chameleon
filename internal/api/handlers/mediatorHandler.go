package handlers

import (
	"github.com/yukitsune/chameleon/pkg/ioc"
	"github.com/yukitsune/chameleon/pkg/mediator"
)

type MediatorHandler struct {
	container ioc.Container
}

func NewMediatorHandler(c ioc.Container) *MediatorHandler {
	return &MediatorHandler{
		container: c,
	}
}

func (h *MediatorHandler) Handle(v interface{}) (interface{}, error) {
	return h.container.ResolveInScopeWithResponse(func (m *mediator.Mediator) (interface{}, error) {
		return m.Send(v)
	})
}
