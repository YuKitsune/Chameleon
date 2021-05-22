package mocks

type MockServiceInterface interface {
	GetReceivedRequest() *MockEvent
}

type MockService struct {
	ReceivedRequest *MockEvent
}

func NewMockService() *MockService {
	return &MockService{}
}

func (m *MockService) DoAThing(r *MockEvent) error {
	m.ReceivedRequest = r
	return nil
}

func (m *MockService) GetReceivedRequest() *MockEvent {
	return m.ReceivedRequest
}

type MockService2ElectricBoogaloo struct {
	ReceivedRequest *MockEvent
}

func NewMockService2ElectricBoogaloo() *MockService2ElectricBoogaloo {
	return &MockService2ElectricBoogaloo{}
}

func (m *MockService2ElectricBoogaloo) DoAThing(r *MockEvent) error {
	m.ReceivedRequest = r
	return nil
}

func (m *MockService2ElectricBoogaloo) GetReceivedRequest() *MockEvent {
	return m.ReceivedRequest
}