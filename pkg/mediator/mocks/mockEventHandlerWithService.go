package mocks

type MockEventHandlerWithService struct {
	service *MockService
}

func NewMockHandlerWithService(svc *MockService) (*MockEventHandlerWithService, error) {
	return &MockEventHandlerWithService{service: svc}, nil
}

func (h *MockEventHandlerWithService) Handle(r *MockEvent) error {
	return h.service.DoAThingWithAnEvent(r)
}

type MockEventHandlerWithService2ElectricBoogaloo struct {
	service *MockService2ElectricBoogaloo
}

func NewMockHandlerWithService2ElectricBoogaloo(svc *MockService2ElectricBoogaloo) (*MockEventHandlerWithService2ElectricBoogaloo, error) {
	return &MockEventHandlerWithService2ElectricBoogaloo{service: svc}, nil
}

func (h *MockEventHandlerWithService2ElectricBoogaloo) Handle(r *MockEvent) error {
	return h.service.DoAThingWithAnEvent(r)
}