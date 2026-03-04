package response_handler

import "github.com/gobeetle/fetch/internal/iface"

// WithErrorHandler sets a custom error handler
func (r *ResponseHandler[Rsp]) WithErrorHandler(handler iface.ResponseErrorHandler) *ResponseHandler[Rsp] {
	r.ResponseErrorHandler = handler
	return r
}

// WithCodeHandler sets a custom status code handler
func (r *ResponseHandler[Rsp]) WithCodeHandler(handler iface.ResponseCodeHandler) *ResponseHandler[Rsp] {
	r.ResponseCodeHandler = handler
	return r
}

// WithRawBodyHandler sets a custom raw body handler
func (r *ResponseHandler[Rsp]) WithRawBodyHandler(handler iface.ResponseRawBodyHandler) *ResponseHandler[Rsp] {
	r.ResponseRawBodyHandler = handler
	return r
}

// WithBodyHandler sets a custom body handler
func (r *ResponseHandler[Rsp]) WithBodyHandler(handler iface.ResponseBodyHandler[Rsp]) *ResponseHandler[Rsp] {
	r.ResponseBodyHandler = handler
	return r
}
