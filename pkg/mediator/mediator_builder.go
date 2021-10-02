package mediator

type MediatorBuilder interface {

	// WithResolver specifies the HandlerResolver to resolve handlers from
	WithResolver(HandlerResolver) MediatorBuilder

	// Build creates a new Mediator instance
	Build() Mediator
}

type mediatorBuilder struct {
	resolver HandlerResolver
}

func NewBuilder() MediatorBuilder {
	return &mediatorBuilder{}
}

func (mb *mediatorBuilder) WithResolver(resolver HandlerResolver) MediatorBuilder {
	mb.resolver = resolver
	return mb
}

func (mb *mediatorBuilder) Build() Mediator {
	return &mediator{
		resolver: mb.resolver,
	}
}
