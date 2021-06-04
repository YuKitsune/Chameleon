package ioc

type Container interface {
	RegisterSingletonInstance(interface{}) error
	RegisterSingletonFactory(interface{}) error

	RegisterTransientInstance(interface{}) error
	RegisterTransientFactory(interface{}) error

	RegisterModule(Module) error

	ResolveInScope(interface{}) error
	ResolveInScopeWithResponse(interface{}) (interface{}, error)
}