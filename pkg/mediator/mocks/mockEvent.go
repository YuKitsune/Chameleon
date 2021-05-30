package mocks

type MockEvent struct {
	Value string
}

type EventRecipient interface {
	GetReceivedEvent() *MockEvent
}