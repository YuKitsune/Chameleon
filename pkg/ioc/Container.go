package ioc

type Container interface {
	RegisterSingletonInstance(interface{}) error
	RegisterSingletonFactory(interface{}) error

	RegisterTransientInstance(interface{}) error
	RegisterTransientFactory(interface{}) error

	ResolveInScope(interface{}) error
}