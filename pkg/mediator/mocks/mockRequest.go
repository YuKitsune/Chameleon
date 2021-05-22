package mocks

type MockRequest struct {
	Value string
}

type MockRequestHandler interface {
	GetReceivedRequest() *MockRequest
}