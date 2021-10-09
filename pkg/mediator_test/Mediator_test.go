package mediator_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/pkg/mediator"
	"github.com/yukitsune/chameleon/pkg/mediator_test/mocks"
	"testing"
)

func TestMediator_Publish(t *testing.T) {

	t.Run("No Event Handlers", func(t *testing.T) {

		// Arrange
		// Set up the Container and Mediator
		c := camogo.NewBuilder().Build()
		m := mediator.NewBuilder().WithResolver(c).Build()

		// Act
		// Create and publish the event
		err := m.Publish(&mocks.MockEvent{
			Value: t.Name(),
		})

		// Assert
		// Todo: Kinda don't want event publishing to error out just because nobody was listening.
		assert.Error(t, err)
	})

	t.Run("Single Event Handler", func(t *testing.T) {

		// Arrange
		// Set up the Container
		h := &mocks.MockEventHandler{}

		cb := camogo.NewBuilder()
		err := cb.RegisterInstance(h)
		assert.NoError(t, err)

		c := cb.Build()

		// Set up the Mediator
		m := mediator.NewBuilder().WithResolver(c).Build()

		// Act
		// Create and publish the event
		e := &mocks.MockEvent{Value: t.Name()}
		err = m.Publish(e)

		// Assert
		assert.NoError(t, err)
		assert.Same(t, e, h.ReceivedEvent)
	})

	t.Run("Multiple Event Handlers", func(t *testing.T) {

		// Arrange
		// Set up the Container
		h1 := &mocks.MockEventHandler{}
		h2 := &mocks.MockEventHandler2ElectricBoogaloo{}

		cb := camogo.NewBuilder()
		err := cb.RegisterInstance(h1)
		assert.NoError(t, err)

		err = cb.RegisterInstance(h2)
		assert.NoError(t, err)

		c := cb.Build()

		// Set up the Mediator
		m := mediator.NewBuilder().WithResolver(c).Build()

		// Act
		// Create and publish the event
		e := &mocks.MockEvent{Value: t.Name()}
		err = m.Publish(e)

		// Assert
		assert.NoError(t, err)
		assert.Same(t, e, h1.ReceivedEvent)
		assert.Same(t, e, h2.ReceivedEvent)
	})
}

func TestMediator_Send(t *testing.T) {

	t.Run("With no Request Handlers", func(t *testing.T) {

		// Arrange
		// Set up the Container and Mediator
		c := camogo.NewBuilder().Build()
		m := mediator.NewBuilder().WithResolver(c).Build()

		// Act
		// Create and send the request
		res, err := m.Send(&mocks.MockRequest{
			Value: t.Name(),
		})

		// Assert
		assert.Nil(t, res)
		assert.Error(t, err)
	})

	t.Run("With Request Handler", func(t *testing.T) {

		// Arrange
		// Set up the Container
		h := &mocks.MockRequestHandler{}

		cb := camogo.NewBuilder()
		err := cb.RegisterInstance(h)
		assert.NoError(t, err)

		c := cb.Build()

		// Set up the Mediator
		m := mediator.NewBuilder().WithResolver(c).Build()

		// Act
		// Create and send the request
		req := &mocks.MockRequest{Value: t.Name()}
		res, err  := m.Send(req)

		// Assert
		assert.NoError(t, err)
		assert.Same(t, req, h.ReceivedRequest)

		typedRes := res.(*mocks.MockResponse)
		assert.Equal(t, req.Value, typedRes.Value)
	})
}
