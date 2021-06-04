package ioc

type Module interface {
	Register(Container) error
}