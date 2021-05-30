package mocks

type MockEventHandler struct {
	ReceivedEvent *MockEvent
}

func (h *MockEventHandler) Handle(e *MockEvent) error {
	h.ReceivedEvent = e
	return nil
}

func (h *MockEventHandler) GetReceivedEvent() *MockEvent {
	return h.ReceivedEvent
}

// Need a second one so we can register two at once without a collision

type MockEventHandler2ElectricBoogaloo struct {
	ReceivedEvent *MockEvent
}

func (h *MockEventHandler2ElectricBoogaloo) Handle(e *MockEvent) error {
	h.ReceivedEvent = e
	return nil
}

func (h *MockEventHandler2ElectricBoogaloo) GetReceivedEvent() *MockEvent {
	return h.ReceivedEvent
}