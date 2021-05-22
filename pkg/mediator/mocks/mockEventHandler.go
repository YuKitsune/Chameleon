package mocks

type MockNormalEventHandler struct {
	ReceivedEvent *MockEvent
}

func (h *MockNormalEventHandler) Handle(e *MockEvent) error {
	h.ReceivedEvent = e
	return nil
}

func (h *MockNormalEventHandler) GetReceivedEvent() *MockEvent {
	return h.ReceivedEvent
}

type MockNormalEventHandler2ElectricBoogaloo struct {
	ReceivedEvent *MockEvent
}

func (h *MockNormalEventHandler2ElectricBoogaloo) Handle(e *MockEvent) error {
	h.ReceivedEvent = e
	return nil
}

func (h *MockNormalEventHandler2ElectricBoogaloo) GetReceivedEvent() *MockEvent {
	return h.ReceivedEvent
}