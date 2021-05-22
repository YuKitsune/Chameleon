package mediator

import (
	"fmt"
	"github.com/yukitsune/chameleon/pkg/ioc"
	"github.com/yukitsune/chameleon/pkg/mediator/mocks"
	"reflect"
	"strings"
	"testing"
)

func TestMediator_Publish(t *testing.T) {

	t.Run("No Event Handlers", func (t *testing.T) {
		c := ioc.NewGolobbyContainer()
		m := New(c)

		r := mocks.MockEvent{
			Value: t.Name(),
		}

		err := m.Publish(r)
		if err != nil {
			failOnError(t, err, "Send() should not return an error")
			return
		}
	})

	t.Run("Single Event Handler Instance", func (t *testing.T) {
		testMediatorPublishWithInstance([]interface{} { &mocks.MockNormalEventHandler{} }, t)
	})

	t.Run("Multiple Event Handler Instances", func (t *testing.T) {
		testMediatorPublishWithInstance(
			[]interface{} {
				&mocks.MockNormalEventHandler{},
				&mocks.MockNormalEventHandler2ElectricBoogaloo{},
			},
			t)
	})

	t.Run("Single Event Handler Factory", func (t *testing.T) {
		testMediatorPublishWithFactory(
			[]interface{} { mocks.NewMockHandlerWithService },
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

	t.Run("Multiple Event Handler Factories", func (t *testing.T) {
		testMediatorPublishWithFactory(
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

func testMediatorPublishWithInstance(handlers []interface{}, t *testing.T) {

	m, _, err := setupMediator(
		&handlers,
		nil,
		nil,
		nil,
		nil)
	if err != nil {
		failOnError(t, err, "setting up mediator should not fail")
		return
	}

	r := &mocks.MockEvent{
		Value: t.Name(),
	}

	err = m.Publish(r)
	if err != nil {
		failOnError(t, err, "Publish() should not return an error")
		return
	}

	for _, requestHandler := range handlers {
		h := requestHandler.(mocks.MockEventHandler)
		receivedRequest := h.GetReceivedEvent()
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

func testMediatorPublishWithFactory(
	factories []interface{},
	serviceProvider func(container ioc.Container) ([]mocks.MockServiceInterface, error),
	t *testing.T) {
	var err error

	m, services, err := setupMediator(
		nil,
		&factories,
		nil,
		nil,
		&serviceProvider)
	if err != nil {
		failOnError(t, err, "setting up mediator should not fail")
	}

	r := &mocks.MockEvent{
		Value: t.Name(),
	}

	err = m.Publish(r)
	if err != nil {
		failOnError(t, err, "Send() should not return an error")
		return
	}

	for _, service := range services {
		receivedRequest := service.GetReceivedEvent()
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

func TestMediator_Send(t *testing.T) {

	t.Run("No Request Handlers", func (t *testing.T) {
		c := ioc.NewGolobbyContainer()
		m := New(c)

		r := mocks.MockRequest{
			Value: t.Name(),
		}

		res, err := m.Send(r)
		if res != nil {
			t.Logf("Expected nil, found %T", res)
			t.Fail()
			return
		}

		if err != nil {
			failOnError(t, err, "Send() should not return an error")
			return
		}
	})

	t.Run("Request Handler Instance", func (t *testing.T) {
		testMediatorSendWithInstance(&mocks.MockNormalRequestHandler{}, t)
	})

	t.Run("Request Handler Factory", func (t *testing.T) {
		testMediatorSendWithFactory(
			mocks.NewMockRequestHandlerWithService,
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
}

func testMediatorSendWithInstance(handler interface{}, t *testing.T) {

	m, _, err := setupMediator(
		nil,
		nil,
		&[]interface{} { handler },
		nil,
		nil)
	if err != nil {
		failOnError(t, err, "setting up mediator should not fail")
		return
	}

	r := &mocks.MockRequest{
		Value: t.Name(),
	}

	res, err := m.Send(r)
	if err != nil {
		failOnError(t, err, "Send() should not return an error")
		return
	}

	if res == nil {
		t.Logf("Expected *mocks.MockResponse, found nil")
		t.Fail()
		return
	}

	response := res.(*mocks.MockResponse)
	if response.Value != r.Value {
		t.Logf("Expected %s, found %s", r.Value, response.Value)
		t.Fail()
	}
}

func testMediatorSendWithFactory(
	factory interface{},
	serviceProvider func(container ioc.Container) ([]mocks.MockServiceInterface, error),
	t *testing.T) {
	var err error

	m, _, err := setupMediator(
		nil,
		nil,
		nil,
		&[]interface{} { factory },
		&serviceProvider)
	if err != nil {
		failOnError(t, err, "setting up mediator should not fail")
		return
	}

	r := &mocks.MockRequest{
		Value: t.Name(),
	}

	res, err := m.Send(r)
	if err != nil {
		failOnError(t, err, "Send() should not return an error")
		return
	}

	if res == nil {
		t.Logf("Expected *mocks.MockResponse, found nil")
		t.Fail()
		return
	}

	response := res.(*mocks.MockResponse)
	if response.Value != r.Value {
		t.Logf("Expected %s, found %s", r.Value, response.Value)
		t.Fail()
	}
}

func TestMediator_Publish_ReturnsUsefulError(t *testing.T) {

	t.Run("Single Event Handler Instance", func (t *testing.T) {
		instances := &[]interface {} {
			&mocks.MockEventHandlerThatAlwaysFails{},
		}
		m, _, err := setupMediator(
			instances,
			nil,
			nil,
			nil,
			nil)
		if err != nil {
			failOnError(t, err, "setting up mediator should not fail")
			return
		}

		r := &mocks.MockEvent{
			Value: t.Name(),
		}
		sendErr := m.Publish(r)
		testErrorIsUseful(sendErr, instances, nil, t)
	})

	t.Run("Multiple Event Handler Instances", func (t *testing.T) {
		instances := &[]interface {} {
			&mocks.MockEventHandlerThatAlwaysFails{},
			&mocks.MockEventHandlerThatAlwaysFails2ElectricBoogaloo{},
		}
		m, _, err := setupMediator(
			instances,
			nil,
			nil,
			nil,
			nil)
		if err != nil {
			failOnError(t, err, "setting up mediator should not fail")
			return
		}

		r := &mocks.MockEvent{
			Value: t.Name(),
		}
		sendErr := m.Publish(r)
		testErrorIsUseful(sendErr, instances, nil, t)
	})

	t.Run("Single Event Handler Factory", func (t *testing.T) {
		factories := &[]interface{} {
			func() *mocks.MockEventHandlerThatAlwaysFails { return &mocks.MockEventHandlerThatAlwaysFails{} },
		}
		m, _, err := setupMediator(
			nil,
			factories,
			nil,
			nil,
			nil)
		if err != nil {
			failOnError(t, err, "setting up mediator should not fail")
			return
		}

		r := &mocks.MockEvent{
			Value: t.Name(),
		}
		sendErr := m.Publish(r)
		testErrorIsUseful(sendErr, nil, factories, t)
	})

	t.Run("Multiple Event Handler Factories", func (t *testing.T) {
		factories := &[]interface{} {
			func() *mocks.MockEventHandlerThatAlwaysFails { return &mocks.MockEventHandlerThatAlwaysFails{} },
			func() *mocks.MockEventHandlerThatAlwaysFails2ElectricBoogaloo { return &mocks.MockEventHandlerThatAlwaysFails2ElectricBoogaloo{} },
		}
		m, _, err := setupMediator(
			nil,
			factories,
			nil,
			nil,
			nil)
		if err != nil {
			failOnError(t, err, "setting up mediator should not fail")
			return
		}

		r := &mocks.MockEvent{
			Value: t.Name(),
		}
		sendErr := m.Publish(r)
		testErrorIsUseful(sendErr, nil, factories, t)
	})
}
func TestMediator_Send_ReturnsUsefulError(t *testing.T) {

	t.Run("Request Handler Instance", func (t *testing.T) {
		instance := &mocks.MockRequestHandlerThatAlwaysFails{}
		m, _, err := setupMediator(
			nil,
			nil,
			&[]interface{} { instance },
			nil,
			nil)
		if err != nil {
			failOnError(t, err, "setting up mediator should not fail")
			return
		}

		r := &mocks.MockRequest{
			Value: t.Name(),
		}
		res, sendErr := m.Send(r)
		testErrorIsUseful(sendErr, &[]interface{} { instance }, nil, t)
		if res != nil {
			t.Logf("Expected nil response, found %v", res)
			t.Fail()
			return
		}
	})

	t.Run("Request Handler Factory", func (t *testing.T) {
		factory := func() *mocks.MockRequestHandlerThatAlwaysFails { return &mocks.MockRequestHandlerThatAlwaysFails{} }
		m, _, err := setupMediator(
			nil,
			nil,
			nil,
			&[]interface{} { factory },
			nil)
		if err != nil {
			failOnError(t, err, "setting up mediator should not fail")
			return
		}

		r := &mocks.MockRequest{
			Value: t.Name(),
		}
		res, sendErr := m.Send(r)
		testErrorIsUseful(sendErr, nil, &[]interface{} { factory }, t)
		if res != nil {
			t.Logf("Expected nil response, found %v", res)
			t.Fail()
			return
		}
	})
}

func setupMediator(
	eventHandlerInstances *[]interface{},
	eventHandlerFactories *[]interface{},
	requestHandlerInstances *[]interface{},
	requestHandlerFactories *[]interface{},
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

	if eventHandlerInstances != nil {
		hs := *eventHandlerInstances
		for _, h := range hs {
			err = m.AddEventHandlerInstance(h)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	if eventHandlerFactories != nil {
		hs := *eventHandlerFactories
		for _, h := range hs {
			err = m.AddEventHandlerFactory(h)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	if requestHandlerInstances != nil {
		hs := *requestHandlerInstances
		for _, h := range hs {
			err = m.AddRequestHandlerInstance(h)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	if requestHandlerFactories != nil {
		hs := *requestHandlerFactories
		for _, h := range hs {
			err = m.AddRequestHandlerFactory(h)
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
