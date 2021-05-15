package mediator

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

type TestRequestHandler interface {
	GetReceivedRequest() *TestRequest
}

type TestRequest struct {
	Value string
}

type TestHandler struct {
	ReceivedRequest *TestRequest
}

func (h *TestHandler) Handle(r *TestRequest) error {
	h.ReceivedRequest = r
	return nil
}

func (h *TestHandler) GetReceivedRequest() *TestRequest {
	return h.ReceivedRequest
}

type TestHandler2ElectricBoogaloo struct {
	ReceivedRequest *TestRequest
}

func (h *TestHandler2ElectricBoogaloo) Handle(r *TestRequest) error {
	h.ReceivedRequest = r
	return nil
}

func (h *TestHandler2ElectricBoogaloo) GetReceivedRequest() *TestRequest {
	return h.ReceivedRequest
}

type TestHandlerThatAlwaysFails struct {
}

func (h *TestHandlerThatAlwaysFails) Handle(r *TestRequest) error {
	return errors.New("oh noes!!!, i did a whoopsie doopsie")
}

type TestHandlerThatAlwaysFails2ElectricBoogaloo struct {
}

func (h *TestHandlerThatAlwaysFails2ElectricBoogaloo) Handle(r *TestRequest) error {
	return errors.New("i can't think of a better error message. it's midnight, i wanna go to bed")
}

func TestMediator_Send_NoHandlers(t *testing.T) {
	m := New()

	testValue := "TestMediator_Send"
	r := &TestRequest{
		Value: testValue,
	}

	err := m.Send(r)
	if err != nil {
		t.Logf("Send() should not return an error, found %v", err)
		t.Fail()
	}
}

func TestMediator_Send_SingleHandler(t *testing.T) {
	h := &TestHandler{}
	m := New()
	err := m.AddHandler(h)
	if err != nil {
		t.Logf("AddHandler() should not return an error, found %v", err)
		t.Fail()
		return
	}

	testValue := "TestMediator_Send"
	r := &TestRequest{
		Value: testValue,
	}

	err = m.Send(r)
	if err != nil {
		t.Logf("Send() should not return an error, found %v", err)
		t.Fail()
		return
	}

	if h.ReceivedRequest != r {
		t.Logf("Handler did not received expected request, expected %v, found %v", r, h.ReceivedRequest)
		t.Fail()
		return
	}
}

func TestMediator_Send_MultipleHandlers(t *testing.T) {
	var handlers []TestRequestHandler
	handlers = append(handlers, &TestHandler{})
	handlers = append(handlers, &TestHandler2ElectricBoogaloo{})

	m := New()
	var err error
	for _, requestHandler := range handlers {
		err = m.AddHandler(requestHandler)
		if err != nil {
			t.Logf("AddHandler() should not return an error, found %v", err)
			t.Fail()
			return
		}
	}

	testValue := "TestMediator_Send"
	r := &TestRequest{
		Value: testValue,
	}

	err = m.Send(r)
	if err != nil {
		t.Logf("Send() should not return an error, found %v", err)
		t.Fail()
		return
	}

	for _, requestHandler := range handlers {
		receivedRequest := requestHandler.GetReceivedRequest()
		if receivedRequest != r {
			t.Logf("Handler did not received expected request, expected %v, found %v", r, receivedRequest)
			t.Fail()
			return
		}
	}
}

func TestMediator_Send_ReturnsError_SingleHandler(t *testing.T) {
	h := &TestHandlerThatAlwaysFails{}
	m := New()
	err := m.AddHandler(h)
	if err != nil {
		t.Logf("AddHandler() should not return an error, found %v", err)
		t.Fail()
		return
	}

	testValue := "TestMediator_Send"
	r := &TestRequest{
		Value: testValue,
	}

	err = m.Send(r)
	if err == nil {
		t.Logf("Send() should return an error, found nil")
		t.Fail()
		return
	}

	if !strings.Contains(err.Error(), fmt.Sprintf("%T", h)) {
		t.Logf("Error must include offending handler")
		t.Fail()
		return
	}
}

func TestMediator_Send_ReturnsError_MultipleHandlers(t *testing.T) {
	var handlers []interface{}
	handlers = append(handlers, &TestHandlerThatAlwaysFails{})
	handlers = append(handlers, &TestHandlerThatAlwaysFails2ElectricBoogaloo{})

	m := New()
	var err error
	for _, requestHandler := range handlers {
		err = m.AddHandler(requestHandler)
		if err != nil {
			t.Logf("AddHandler() should not return an error, found %v", err)
			t.Fail()
			return
		}
	}

	testValue := "TestMediator_Send"
	r := &TestRequest{
		Value: testValue,
	}

	err = m.Send(r)
	if err == nil {
		t.Logf("Send() should return an error, found nil")
		t.Fail()
		return
	}

	for _, requestHandler := range handlers {
		if !strings.Contains(err.Error(), fmt.Sprintf("%T", requestHandler)) {
			t.Logf("Error must include offending handler")
			t.Fail()
			return
		}
	}
}
