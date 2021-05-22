package mocks

type MockNormalRequestHandler struct {
	ReceivedRequest *MockRequest
	Response *MockResponse
}

func (h *MockNormalRequestHandler) Handle(e *MockRequest) (*MockResponse, error) {
	h.ReceivedRequest = e
	h.Response = &MockResponse{Value: e.Value}
	return h.Response, nil
}

func (h *MockNormalRequestHandler) GetReceivedRequest() *MockResponse {
	return h.Response
}