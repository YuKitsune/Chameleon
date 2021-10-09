package mocks

type MockRequestHandler struct {
	ReceivedRequest *MockRequest
	Response        *MockResponse
}

func (h *MockRequestHandler) Handle(e *MockRequest) (*MockResponse, error) {
	h.ReceivedRequest = e
	h.Response = &MockResponse{Value: e.Value}
	return h.Response, nil
}
