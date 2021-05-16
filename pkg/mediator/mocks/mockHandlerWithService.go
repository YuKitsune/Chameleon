package mocks

type MockHandlerWithService struct {
	service *MockService
}

func NewMockHandlerWithService(svc *MockService) (*MockHandlerWithService, error) {
	return &MockHandlerWithService{service: svc}, nil
}

func (h *MockHandlerWithService) Handle(r *MockRequest) error {
	return h.service.DoAThing(r)
}

type MockHandlerWithService2ElectricBoogaloo struct {
	service *MockService2ElectricBoogaloo
}

func NewMockHandlerWithService2ElectricBoogaloo(svc *MockService2ElectricBoogaloo) (*MockHandlerWithService2ElectricBoogaloo, error) {
	return &MockHandlerWithService2ElectricBoogaloo{service: svc}, nil
}

func (h *MockHandlerWithService2ElectricBoogaloo) Handle(r *MockRequest) error {
	return h.service.DoAThing(r)
}