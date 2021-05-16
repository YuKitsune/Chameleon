package mocks

type MockServiceInterface interface {
	GetReceivedRequest() *MockRequest
}

type MockService struct {
	ReceivedRequest *MockRequest
}

func NewMockService() *MockService {
	return &MockService{}
}

func (m *MockService) DoAThing(r *MockRequest) error {
	m.ReceivedRequest = r
	return nil
}

func (m *MockService) GetReceivedRequest() *MockRequest {
	return m.ReceivedRequest
}

type MockService2ElectricBoogaloo struct {
	ReceivedRequest *MockRequest
}

func NewMockService2ElectricBoogaloo() *MockService2ElectricBoogaloo {
	return &MockService2ElectricBoogaloo{}
}

func (m *MockService2ElectricBoogaloo) DoAThing(r *MockRequest) error {
	m.ReceivedRequest = r
	return nil
}

func (m *MockService2ElectricBoogaloo) GetReceivedRequest() *MockRequest {
	return m.ReceivedRequest
}