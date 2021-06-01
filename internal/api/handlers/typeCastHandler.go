package handlers

type castFn func (interface{}) interface{}

type TypeCastHandler struct {
	castRequest *castFn
	castResponse *castFn
	inner Handler
}

func NewTypeCastHandler(castRequest castFn, inner Handler) *TypeCastHandler {
	return &TypeCastHandler{
		castRequest: &castRequest,
		inner: inner,
	}
}

func NewTypeCastHandlerWithResponse(castRequest castFn, castResponse castFn, inner Handler) *TypeCastHandler {
	return &TypeCastHandler{
		castRequest: &castRequest,
		castResponse: &castResponse,
		inner: inner,
	}
}

func (h TypeCastHandler) Handle(v interface{}) (interface{}, error) {

	var err error

	// Cast the request if necessary
	req := v
	if h.castRequest != nil {
		castFn := *h.castRequest
		req = castFn(v)
	}

	// Send the casted request to the next handler
	res, err := h.inner.Handle(req)
	if err != nil {
		return nil, err
	}

	// Cast the response if necessary
	if h.castResponse != nil {
		castFn := *h.castResponse
		res = castFn(res)
	}

	return res, nil
}