package alias

import "github.com/yukitsune/chameleon/pkg/ioc"

type AliasHandlerModule struct {}

func NewAliasHandlerModule() *AliasHandlerModule {
	return &AliasHandlerModule{}
}

func (module *AliasHandlerModule) Register(container ioc.Container) error {
	var err error

	err = container.RegisterTransientFactory(NewCreateAliasHandler)
	if err != nil {
		return err
	}

	err = container.RegisterTransientFactory(NewReadAliasHandler)
	if err != nil {
		return err
	}

	err = container.RegisterTransientFactory(NewUpdateAliasHandler)
	if err != nil {
		return err
	}

	err = container.RegisterTransientFactory(NewDeleteAliasHandler)
	if err != nil {
		return err
	}

	return nil
}