package mocks

import "errors"

type MockEventHandlerThatAlwaysFails struct {
}

func (h *MockEventHandlerThatAlwaysFails) Handle(e *MockEvent) error {
	return errors.New("oh noes!!!, i did a whoopsie doopsie")
}

type MockEventHandlerThatAlwaysFails2ElectricBoogaloo struct {
}

func (h *MockEventHandlerThatAlwaysFails2ElectricBoogaloo) Handle(e *MockEvent) error {
	return errors.New("i can't think of a better error message. it's midnight, i wanna go to bed")
}