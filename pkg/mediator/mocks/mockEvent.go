package mocks

type MockEvent struct {
	Value string
}

type MockEventHandler interface {
	GetReceivedEvent() *MockEvent
}