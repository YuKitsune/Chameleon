package mocks

import "errors"

type MockRequestHandlerThatAlwaysFails struct {
}

func (h *MockRequestHandlerThatAlwaysFails) Handle(r *MockRequest) (*MockResponse, error) {
	return nil, errors.New("oh noes!!!, i did a whoopsie doopsie")
}