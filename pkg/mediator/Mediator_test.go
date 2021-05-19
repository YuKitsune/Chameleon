package mediator

import (
	"fmt"
	"github.com/yukitsune/chameleon/pkg/ioc"
	"github.com/yukitsune/chameleon/pkg/mediator/mocks"
	"reflect"
	"strings"
	"testing"
)

func TestMediator_Send(t *testing.T) {

	t.Run("No Handlers", func (t *testing.T) {
		c := ioc.NewGolobbyContainer()
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
		testMediatorSendWithInstance([]interface{} { &mocks.MockHandler{} }, t)
	})

	t.Run("Multiple Handler Instances", func (t *testing.T) {
		testMediatorSendWithInstance(
			[]interface{} {
				&mocks.MockHandler{},
				&mocks.MockHandler2ElectricBoogaloo{},
			},
			t)
	})

	t.Run("Single Handler Factory", func (t *testing.T) {
		testMediatorSendWithFactory(
			[]interface{} {mocks.NewMockHandlerWithService },
			func (c ioc.Container) ([]mocks.MockServiceInterface, error) {
				service := mocks.NewMockService()
				err := c.RegisterSingletonInstance(service)
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
			func (c ioc.Container) ([]mocks.MockServiceInterface, error) {
				var services []mocks.MockServiceInterface

				firstService := mocks.NewMockService()
				services = append(services, firstService)
				err := c.RegisterSingletonInstance(firstService)
				if err != nil {
					return nil, err
				}

				secondService := mocks.NewMockService2ElectricBoogaloo()
				services = append(services, secondService)
				err = c.RegisterSingletonInstance(secondService)
				if err != nil {
					return nil, err
				}

				return services, nil
			},
			t)
	})
}

func testMediatorSendWithInstance(handlers []interface{}, t *testing.T) {

	m, _, err := setupMediator(
		&handlers,
		nil,
		nil)
	if err != nil {
		failOnError(t, err, "setting up mediator should not fail")
	}

	r := &mocks.MockRequest{
		Value: t.Name(),
	}

	err = m.Send(r)
	if err != nil {
		failOnError(t, err, "Send() should not return an error")
		return
	}

	for _, requestHandler := range handlers {
		h := requestHandler.(mocks.MockRequestHandler)
		receivedRequest := h.GetReceivedRequest()
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
	serviceProvider func(container ioc.Container) ([]mocks.MockServiceInterface, error),
	t *testing.T) {
	var err error

	m, services, err := setupMediator(
		nil,
		&factories,
		&serviceProvider)
	if err != nil {
		failOnError(t, err, "setting up mediator should not fail")
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

func TestMediator_Send_ReturnsUsefulError(t *testing.T) {

	t.Run("Single Handler Instance", func (t *testing.T) {
		instances := &[]interface {} {
			&mocks.MockHandlerThatAlwaysFails{},
		}
		m, _, err := setupMediator(
			instances,
			nil,
			nil)
		if err != nil {
			failOnError(t, err, "setting up mediator should not fail")
			return
		}

		r := &mocks.MockRequest{
			Value: t.Name(),
		}
		sendErr := m.Send(r)
		testErrorIsUseful(sendErr, instances, nil, t)
	})

	t.Run("Multiple Handler Instances", func (t *testing.T) {
		instances := &[]interface {} {
			&mocks.MockHandlerThatAlwaysFails{},
			&mocks.MockHandlerThatAlwaysFails2ElectricBoogaloo{},
		}
		m, _, err := setupMediator(
			instances,
			nil,
			nil)
		if err != nil {
			failOnError(t, err, "setting up mediator should not fail")
			return
		}

		r := &mocks.MockRequest{
			Value: t.Name(),
		}
		sendErr := m.Send(r)
		testErrorIsUseful(sendErr, instances, nil, t)
	})

	t.Run("Single Handler Factory", func (t *testing.T) {
		factories := &[]interface{} {
			func() *mocks.MockHandlerThatAlwaysFails { return &mocks.MockHandlerThatAlwaysFails{} },
		}
		m, _, err := setupMediator(
			nil,
			factories,
			nil)
		if err != nil {
			failOnError(t, err, "setting up mediator should not fail")
			return
		}

		r := &mocks.MockRequest{
			Value: t.Name(),
		}
		sendErr := m.Send(r)
		testErrorIsUseful(sendErr, nil, factories, t)
	})

	t.Run("Multiple Handler Factories", func (t *testing.T) {
		factories := &[]interface{} {
			func() *mocks.MockHandlerThatAlwaysFails { return &mocks.MockHandlerThatAlwaysFails{} },
			func() *mocks.MockHandlerThatAlwaysFails2ElectricBoogaloo { return &mocks.MockHandlerThatAlwaysFails2ElectricBoogaloo{} },
		}
		m, _, err := setupMediator(
			nil,
			factories,
			nil)
		if err != nil {
			failOnError(t, err, "setting up mediator should not fail")
			return
		}

		r := &mocks.MockRequest{
			Value: t.Name(),
		}
		sendErr := m.Send(r)
		testErrorIsUseful(sendErr, nil, factories, t)
	})
}

func setupMediator(
	handlerInstances *[]interface{},
	handlerFactories *[]interface{},
	serviceProvider *func(container ioc.Container) ([]mocks.MockServiceInterface, error),
	) (*Mediator, []mocks.MockServiceInterface, error) {
	var err error

	c := ioc.NewGolobbyContainer()
	var services []mocks.MockServiceInterface
	if serviceProvider != nil {
		serviceProviderFn := *serviceProvider
		services, err = serviceProviderFn(c)
		if err != nil {
			return nil, nil, err
		}
	}

	m := New(c)

	if handlerInstances != nil {
		hs := *handlerInstances
		for _, h := range hs {
			err = m.AddHandlerInstance(h)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	if handlerFactories != nil {
		hs := *handlerFactories
		for _, h := range hs {
			err = m.AddHandlerFactory(h)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	return m, services, nil
}

func testErrorIsUseful(
	err error,
	handlerInstances *[]interface {},
	handlerFactories *[]interface {},
	t *testing.T) {

	var expectedNumberOfErrors int
	var expectedErrorTypes []string

	if handlerInstances != nil {
		hs := *handlerInstances
		expectedNumberOfErrors += len(hs)

		for _, instance := range hs {
			expectedErrorTypes = append(expectedErrorTypes, fmt.Sprintf("%T", instance))
		}
	}

	if handlerFactories != nil {
		hs := *handlerFactories
		expectedNumberOfErrors += len(hs)

		for _, factory := range hs {
			factoryReturnType := reflect.TypeOf(factory).Out(0)
			expectedErrorTypes = append(expectedErrorTypes, fmt.Sprintf("%s", factoryReturnType.String()))
		}
	}

	errorText := "errors"
	if expectedNumberOfErrors == 1 {
		errorText = "error"
	}

	if err == nil {
		t.Logf("Expected %d %s, found none", expectedNumberOfErrors, errorText)
		t.Fail()
		return
	}

	errors := strings.Split(err.Error(), "\n")
	if len(errors) != expectedNumberOfErrors {
		t.Logf("Expected %d %s, found %d", expectedNumberOfErrors, errorText, len(errors))
		t.Fail()
		return
	}

	for _, errorType := range expectedErrorTypes {
		found := false
		for _, s := range errors {
			if strings.Contains(s, errorType) {
				found = true
				break
			}
		}

		if !found {
			t.Logf("Expected to find an error for %s, found none", errorType)
			t.Fail()
			return
		}
	}
}

func failOnError(t *testing.T, err error, reason string) {
	t.Logf("%s, found %v", reason, err)
	t.Fail()
}
