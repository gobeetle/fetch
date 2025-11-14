package response_handler

import "github.com/gobeetle/fetch/internal/iface"

// WithErrorHandler sets a custom error handler
func (r *ResponseHandler) WithErrorHandler(handler iface.ResponseErrorHandler) *ResponseHandler {
	r.ResponseErrorHandler = handler
	return r
}

// WithCodeHandler sets a custom status code handler
func (r *ResponseHandler) WithCodeHandler(handler iface.ResponseCodeHandler) *ResponseHandler {
	r.ResponseCodeHandler = handler
	return r
}

// WithRawBodyHandler sets a custom raw body handler
func (r *ResponseHandler) WithRawBodyHandler(handler iface.ResponseRawBodyHandler) *ResponseHandler {
	r.ResponseRawBodyHandler = handler
	return r
}

// WithBodyHandler sets a custom body handler
func (r *ResponseHandler) WithBodyHandler(handler iface.ResponseBodyHandler) *ResponseHandler {
	r.ResponseBodyHandler = handler
	return r
}
