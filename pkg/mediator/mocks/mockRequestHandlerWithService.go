package mocks

type MockRequestHandlerWithService struct {
	service *MockService
}

func NewMockRequestHandlerWithService(svc *MockService) (*MockRequestHandlerWithService, error) {
	return &MockRequestHandlerWithService{service: svc}, nil
}

func (h *MockRequestHandlerWithService) Handle(r *MockRequest) (*MockResponse, error) {
	err := h.service.DoAThingWithARequest(r)
	if err != nil {
		return nil, err
	}

	return &MockResponse{
		Value: r.Value,
	}, nil
}