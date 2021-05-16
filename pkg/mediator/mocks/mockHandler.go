package mocks

type MockHandler struct {
	ReceivedRequest *MockRequest
}

func (h *MockHandler) Handle(r *MockRequest) error {
	h.ReceivedRequest = r
	return nil
}

func (h *MockHandler) GetReceivedRequest() *MockRequest {
	return h.ReceivedRequest
}

type MockHandler2ElectricBoogaloo struct {
	ReceivedRequest *MockRequest
}

func (h *MockHandler2ElectricBoogaloo) Handle(r *MockRequest) error {
	h.ReceivedRequest = r
	return nil
}

func (h *MockHandler2ElectricBoogaloo) GetReceivedRequest() *MockRequest {
	return h.ReceivedRequest
}