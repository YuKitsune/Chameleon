package mocks

import "errors"

type MockHandlerThatAlwaysFails struct {
}

func (h *MockHandlerThatAlwaysFails) Handle(r *MockRequest) error {
	return errors.New("oh noes!!!, i did a whoopsie doopsie")
}

type MockHandlerThatAlwaysFails2ElectricBoogaloo struct {
}

func (h *MockHandlerThatAlwaysFails2ElectricBoogaloo) Handle(r *MockRequest) error {
	return errors.New("i can't think of a better error message. it's midnight, i wanna go to bed")
}