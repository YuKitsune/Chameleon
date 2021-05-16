package mediator

import (
	"github.com/yukitsune/chameleon/pkg/mediator/mocks"
	"go.uber.org/dig"
	"testing"
)

func TestMediator_Send(t *testing.T) {

	t.Run("No Handlers", func (t *testing.T) {
		c := dig.New()
		m := New(c)

		r := mocks.MockRequest{
			Value: t.Name(),
		}

		err := m.Send(r)
		if err != nil {
			failOnError(t, err, "Send() should not return an error")
			return
		}
	})
	t.Run("Single Handler Instance", func (t *testing.T) {
		testMediatorSendWithInstance([]mocks.MockRequestHandler { &mocks.MockHandler{} }, t)
	})
	t.Run("Multiple Handler Instances", func (t *testing.T) {
		testMediatorSendWithInstance(
			[]mocks.MockRequestHandler {
				&mocks.MockHandler{},
				&mocks.MockHandler2ElectricBoogaloo{},
			},
			t)
	})
	t.Run("Single Handler Factory", func (t *testing.T) {
		// services := []mocks.MockServiceInterface { &mocks.MockService{} }
		testMediatorSendWithFactory(
			[]interface{} {mocks.NewMockHandlerWithService },
			func (c *dig.Container) ([]mocks.MockServiceInterface, error) {
				service := mocks.NewMockService()
				err := c.Provide(func () *mocks.MockService { return service })
				if err != nil {
					return nil, err
				}

				return []mocks.MockServiceInterface { service }, nil
			},
			t)
	})
	t.Run("Multiple Handler Factories", func (t *testing.T) {
		testMediatorSendWithFactory(
			[]interface{} {
				mocks.NewMockHandlerWithService,
				mocks.NewMockHandlerWithService2ElectricBoogaloo,
			},
			func (c *dig.Container) ([]mocks.MockServiceInterface, error) {
				var services []mocks.MockServiceInterface

				firstService := mocks.NewMockService()
				services = append(services, firstService)
				err := c.Provide(func () *mocks.MockService { return firstService })
				if err != nil {
					return nil, err
				}

				secondService := mocks.NewMockService2ElectricBoogaloo()
				services = append(services, secondService)
				err = c.Provide(func () *mocks.MockService2ElectricBoogaloo { return secondService })
				if err != nil {
					return nil, err
				}

				return services, nil
			},
			t)
	})
}

func testMediatorSendWithInstance(handlers []mocks.MockRequestHandler, t *testing.T) {

	c := dig.New()
	m := New(c)

	for _, requestHandler := range handlers {
		err := m.AddHandlerInstance(requestHandler)
		if err != nil {
			failOnError(t, err, "AddHandlerInstance() should not return an error")
			return
		}
	}

	r := &mocks.MockRequest{
		Value: t.Name(),
	}

	err := m.Send(r)
	if err != nil {
		failOnError(t, err, "Send() should not return an error")
		return
	}

	for _, requestHandler := range handlers {
		receivedRequest := requestHandler.GetReceivedRequest()
		if receivedRequest == nil {
			t.Logf("Expected %T, found nil", r)
			t.Fail()
			return
		}

		if receivedRequest.Value != r.Value {
			t.Logf("Expected %s, found %s", r.Value, receivedRequest)
			t.Fail()
		}
	}
}

func testMediatorSendWithFactory(
	factories []interface{},
	serviceProvider func(container *dig.Container) ([]mocks.MockServiceInterface, error),
	t *testing.T) {
	var err error

	c := dig.New()
	services, err := serviceProvider(c)
	if err != nil {
		failOnError(t, err, "the service provider should not return an error")
		return
	}

	m := New(c)

	for _, factory := range factories {
		err = m.AddHandlerFactory(factory)
		if err != nil {
			failOnError(t, err, "AddHandlerFactory() should not return an error")
			return
		}
	}

	r := &mocks.MockRequest{
		Value: t.Name(),
	}

	err = m.Send(r)
	if err != nil {
		failOnError(t, err, "Send() should not return an error")
		return
	}

	for _, service := range services {
		receivedRequest := service.GetReceivedRequest()
		if receivedRequest == nil {
			t.Logf("Expected %T, found nil", r)
			t.Fail()
			return
		}

		if receivedRequest.Value != r.Value {
			t.Logf("Expected %s, found %s", r.Value, receivedRequest)
			t.Fail()
		}
	}
}

//func TestMediator_Send_ReturnsUsefulError(t *testing.T) {
//	t.Run("Single Handler Instance")
//	t.Run("Multiple Handler Instances")
//	t.Run("Single Handler Factory")
//	t.Run("Multiple Handler Factories")
//}

func failOnError(t *testing.T, err error, reason string) {
	t.Logf("%s, found %v", reason, err)
	t.Fail()
}
