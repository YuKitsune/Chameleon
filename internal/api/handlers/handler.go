package handlers

type Handler interface {
	Handle(interface{}) error
}
