package modules

import (
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/internal/api/handlers/alias"
	"github.com/yukitsune/chameleon/pkg/mediator"
)

type MediatorHandlerModule struct {
}

func (m *MediatorHandlerModule) Register(builder camogo.ContainerBuilder) (err error) {

	err = builder.RegisterFactory(func (container camogo.Container) mediator.Mediator{
		return mediator.NewBuilder().WithResolver(container).Build()
	}, camogo.TransientLifetime)
	if err != nil {
		return err
	}

	err = builder.RegisterFactory(alias.NewCreateAliasHandler, camogo.TransientLifetime)
	if err != nil {
		return err
	}

	err = builder.RegisterFactory(alias.NewReadAliasHandler, camogo.TransientLifetime)
	if err != nil {
		return err
	}

	err = builder.RegisterFactory(alias.NewUpdateAliasHandler, camogo.TransientLifetime)
	if err != nil {
		return err
	}

	err = builder.RegisterFactory(alias.NewDeleteAliasHandler, camogo.TransientLifetime)
	if err != nil {
		return err
	}

	return nil
}
