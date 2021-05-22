package mocks

type MockServiceInterface interface {
	GetReceivedEvent() *MockEvent
	GetReceivedRequest() *MockRequest
}

type MockService struct {
	ReceivedEvent *MockEvent
	ReceivedRequest *MockRequest
}

func NewMockService() *MockService {
	return &MockService{}
}

func (m *MockService) DoAThingWithAnEvent(r *MockEvent) error {
	m.ReceivedEvent = r
	return nil
}

func (m *MockService) DoAThingWithARequest(r *MockRequest) error {
	m.ReceivedRequest = r
	return nil
}

func (m *MockService) GetReceivedEvent() *MockEvent {
	return m.ReceivedEvent
}

func (m *MockService) GetReceivedRequest() *MockRequest {
	return m.ReceivedRequest
}

type MockService2ElectricBoogaloo struct {
	ReceivedEvent *MockEvent
	ReceivedRequest *MockRequest
}

func NewMockService2ElectricBoogaloo() *MockService2ElectricBoogaloo {
	return &MockService2ElectricBoogaloo{}
}

func (m *MockService2ElectricBoogaloo) DoAThingWithAnEvent(r *MockEvent) error {
	m.ReceivedEvent = r
	return nil
}

func (m *MockService2ElectricBoogaloo) DoAThingWithARequest(r *MockRequest) error {
	m.ReceivedRequest = r
	return nil
}

func (m *MockService2ElectricBoogaloo) GetReceivedEvent() *MockEvent {
	return m.ReceivedEvent
}

func (m *MockService2ElectricBoogaloo) GetReceivedRequest() *MockRequest {
	return m.ReceivedRequest
}