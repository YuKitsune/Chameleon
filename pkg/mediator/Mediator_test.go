package mediator

import (
	"fmt"
	"github.com/yukitsune/camogo"
	"github.com/yukitsune/chameleon/pkg/mediator/mocks"
	"github.com/yukitsune/chameleon/pkg/testUtils"
	"reflect"
	"strings"
	"testing"
)

func TestMediator_Publish(t *testing.T) {

	t.Run("No Event Handlers", func(t *testing.T) {
		c := camogo.New()
		m := New(c)

		r := mocks.MockEvent{
			Value: t.Name(),
		}

		err := m.Publish(r)
		if err != nil {
			testUtils.FailOnError(t, err, "Publish() should not return an error")
			return
		}
	})

	t.Run("Single Event Handler Instance", func(t *testing.T) {
		testMediatorPublish(
			nil,
			func(m *Mediator, _ camogo.Container) (*[]mocks.EventRecipient, error) {
				h := &mocks.MockEventHandler{}
				m.AddEventHandlerInstance(h)
				return &[]mocks.EventRecipient{h}, nil
			},
			t)
	})

	t.Run("Multiple Event Handler Instances", func(t *testing.T) {
		testMediatorPublish(
			nil,
			func(m *Mediator, _ camogo.Container) (*[]mocks.EventRecipient, error) {
				h1 := &mocks.MockEventHandler{}
				m.AddEventHandlerInstance(h1)

				h2 := &mocks.MockEventHandler2ElectricBoogaloo{}
				m.AddEventHandlerInstance(h2)

				return &[]mocks.EventRecipient{h1, h2}, nil
			},
			t)
	})

	t.Run("Single Event Handler Factory", func(t *testing.T) {
		setupContainer := func(c camogo.Container) (*[]mocks.EventRecipient, error) {
			err := c.Register(func (r *camogo.Registrar) error {
				return r.RegisterFactory(mocks.NewMockService, camogo.SingletonLifetime)
			})

			if err != nil {
				return nil, err
			}

			// Rip the recipient out of the container so we can assert that it received the event
			var r mocks.EventRecipient
			err = c.Resolve(func(s *mocks.MockService) { r = s })
			if err != nil {
				return nil, err
			}

			return &[]mocks.EventRecipient{r}, nil
		}

		testMediatorPublish(
			&setupContainer,
			func(m *Mediator, c camogo.Container) (*[]mocks.EventRecipient, error) {
				m.AddEventHandlerFactory(mocks.NewMockHandlerWithService)
				return nil, nil
			},
			t)
	})

	t.Run("Multiple Event Handler Factories", func(t *testing.T) {
		setupContainer := func(c camogo.Container) (*[]mocks.EventRecipient, error) {
			err := c.Register(func (r *camogo.Registrar) error {
				err := r.RegisterFactory(mocks.NewMockService, camogo.SingletonLifetime)
				if err != nil {
					return err
				}

				err = r.RegisterFactory(mocks.NewMockService2ElectricBoogaloo, camogo.SingletonLifetime)
				if err != nil {
					return err
				}

				return nil
			})

			if err != nil {
				return nil, err
			}

			// Rip the recipients out of the container so we can assert that it received the event
			var r1 mocks.EventRecipient
			var r2 mocks.EventRecipient
			err = c.Resolve(func(s1 *mocks.MockService, s2 *mocks.MockService2ElectricBoogaloo) {
				r1 = s1
				r2 = s2
			})
			if err != nil {
				return nil, err
			}

			return &[]mocks.EventRecipient{r1, r2}, nil
		}

		testMediatorPublish(
			&setupContainer,
			func(m *Mediator, c camogo.Container) (*[]mocks.EventRecipient, error) {
				m.AddEventHandlerFactory(mocks.NewMockHandlerWithService)
				m.AddEventHandlerFactory(mocks.NewMockHandlerWithService2ElectricBoogaloo)
				return nil, nil
			},
			t)
	})
}

func testMediatorPublish(
	setupContainer *func(camogo.Container) (*[]mocks.EventRecipient, error),
	setupMediator func(*Mediator, camogo.Container) (*[]mocks.EventRecipient, error),
	t *testing.T) {
	var err error

	c := camogo.New()

	// Setup the container if we've been given a function for it
	var recipientsFromServices *[]mocks.EventRecipient
	if setupContainer != nil {
		setupServicesFn := *setupContainer
		recipientsFromServices, err = setupServicesFn(c)
		if err != nil {
			testUtils.FailOnError(t, err, "container setup should not fail")
			return
		}
	}

	m := New(c)
	recipientsFromMediator, err := setupMediator(m, c)
	if err != nil {
		testUtils.FailOnError(t, err, "mediator setup should not fail")
		return
	}

	// Merge all recipients into one slice
	var recipients []mocks.EventRecipient
	if recipientsFromServices != nil && len(*recipientsFromServices) > 0 {
		for _, recipient := range *recipientsFromServices {
			recipients = append(recipients, recipient)
		}
	}

	if recipientsFromMediator != nil && len(*recipientsFromMediator) > 0 {
		for _, recipient := range *recipientsFromMediator {
			recipients = append(recipients, recipient)
		}
	}

	// Publish an event
	r := &mocks.MockEvent{
		Value: t.Name(),
	}

	err = m.Publish(r)
	if err != nil {
		testUtils.FailOnError(t, err, "Publish() should not return an error")
		return
	}

	// Ensure each registered recipient has received the event
	for _, recipient := range recipients {
		receivedRequest := recipient.GetReceivedEvent()
		if receivedRequest == nil {
			t.Logf("Expected %T, found nil", r)
			t.Fail()
			return
		}

		if receivedRequest.Value != r.Value {
			t.Logf("Expected %s, found %s", r.Value, receivedRequest)
			t.Fail()
			return
		}
	}
}

func TestMediator_Send(t *testing.T) {

	t.Run("No Request Handlers", func(t *testing.T) {
		c := camogo.New()
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
			testUtils.FailOnError(t, err, "Send() should not return an error")
			return
		}
	})

	t.Run("Request Handler Instance", func(t *testing.T) {
		testMediatorSend(
			nil,
			func(m *Mediator) error {
				m.AddRequestHandlerInstance(&mocks.MockRequestHandler{})
				return nil
			},
			t)
	})

	t.Run("Request Handler Factory", func(t *testing.T) {
		setupContainer := func(c camogo.Container) error {
			service := mocks.NewMockService()
			return c.Register(func (r *camogo.Registrar) error {
				return r.RegisterInstance(service)
			})
		}

		testMediatorSend(
			&setupContainer,
			func(m *Mediator) error {
				m.AddRequestHandlerFactory(mocks.NewMockRequestHandlerWithService)
				return nil
			},
			t)
	})
}

func testMediatorSend(
	setupContainer *func(camogo.Container) error,
	setupMediator func(*Mediator) error,
	t *testing.T) {
	var err error

	c := camogo.New()

	// Setup the container if we've been given a function for it
	if setupContainer != nil {
		setupServicesFn := *setupContainer
		err = setupServicesFn(c)
		if err != nil {
			testUtils.FailOnError(t, err, "container setup should not fail")
			return
		}
	}

	m := New(c)
	err = setupMediator(m)
	if err != nil {
		testUtils.FailOnError(t, err, "mediator setup should not fail")
		return
	}

	// Send the request
	r := &mocks.MockRequest{
		Value: t.Name(),
	}

	res, err := m.Send(r)
	if err != nil {
		testUtils.FailOnError(t, err, "Send() should not return an error")
		return
	}

	// Ensure the response received matches what was expected
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

	t.Run("Single Event Handler Instance", func(t *testing.T) {
		testMediatorPublishReturnsUsefulError(func(m *Mediator) (*[]interface{}, *[]interface{}) {
			instance := &mocks.MockEventHandlerThatAlwaysFails{}
			m.AddEventHandlerInstance(instance)
			return &[]interface{}{instance}, nil
		},
			t)
	})

	t.Run("Multiple Event Handler Instances", func(t *testing.T) {
		testMediatorPublishReturnsUsefulError(func(m *Mediator) (*[]interface{}, *[]interface{}) {
			instances := &[]interface{}{
				&mocks.MockEventHandlerThatAlwaysFails{},
				&mocks.MockEventHandlerThatAlwaysFails2ElectricBoogaloo{},
			}

			for _, instance := range *instances {
				m.AddEventHandlerInstance(instance)
			}

			return instances, nil
		},
			t)
	})

	t.Run("Single Event Handler Factory", func(t *testing.T) {
		testMediatorPublishReturnsUsefulError(func(m *Mediator) (*[]interface{}, *[]interface{}) {
			factory := func() *mocks.MockEventHandlerThatAlwaysFails { return &mocks.MockEventHandlerThatAlwaysFails{} }
			m.AddEventHandlerFactory(factory)
			return nil, &[]interface{}{factory}
		},
			t)
	})

	t.Run("Multiple Event Handler Factories", func(t *testing.T) {
		testMediatorPublishReturnsUsefulError(func(m *Mediator) (*[]interface{}, *[]interface{}) {
			factories := &[]interface{}{
				func() *mocks.MockEventHandlerThatAlwaysFails { return &mocks.MockEventHandlerThatAlwaysFails{} },
				func() *mocks.MockEventHandlerThatAlwaysFails2ElectricBoogaloo {
					return &mocks.MockEventHandlerThatAlwaysFails2ElectricBoogaloo{}
				},
			}

			for _, factory := range *factories {
				m.AddEventHandlerFactory(factory)
			}

			return nil, factories
		},
			t)
	})
}

func testMediatorPublishReturnsUsefulError(setupMediator func(*Mediator) (*[]interface{}, *[]interface{}), t *testing.T) {
	c := camogo.New()
	m := New(c)

	handlerInstances, handlerFactories := setupMediator(m)

	r := &mocks.MockEvent{
		Value: t.Name(),
	}
	pubErr := m.Publish(r)
	testErrorIsUseful(pubErr, handlerInstances, handlerFactories, t)
}

func TestMediator_Send_ReturnsUsefulError(t *testing.T) {

	t.Run("Request Handler Instance", func(t *testing.T) {
		testMediatorSendReturnsUsefulError(func(m *Mediator) (*[]interface{}, *[]interface{}) {
			instance := &mocks.MockRequestHandlerThatAlwaysFails{}
			m.AddRequestHandlerInstance(instance)
			return &[]interface{}{instance}, nil
		},
			t)
	})

	t.Run("Request Handler Factory", func(t *testing.T) {
		testMediatorSendReturnsUsefulError(func(m *Mediator) (*[]interface{}, *[]interface{}) {
			factory := func() *mocks.MockRequestHandlerThatAlwaysFails { return &mocks.MockRequestHandlerThatAlwaysFails{} }
			m.AddRequestHandlerFactory(factory)
			return nil, &[]interface{}{factory}
		},
			t)
	})
}

func testMediatorSendReturnsUsefulError(setupMediator func(*Mediator) (*[]interface{}, *[]interface{}), t *testing.T) {
	c := camogo.New()
	m := New(c)

	handlerInstances, handlerFactories := setupMediator(m)

	r := &mocks.MockRequest{
		Value: t.Name(),
	}
	res, sendErr := m.Send(r)
	testErrorIsUseful(sendErr, handlerInstances, handlerFactories, t)
	if res != nil {
		t.Logf("Expected nil response, found %v", res)
		t.Fail()
		return
	}
}

func testErrorIsUseful(
	err error,
	handlerInstances *[]interface{},
	handlerFactories *[]interface{},
	t *testing.T) {

	// Figure out how many errors we should have and collate a list of types that should've returned errors
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