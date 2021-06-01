package handlers

type Handler interface {
	Handle(interface{}) (interface{}, error)
}
