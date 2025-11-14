package request_handler

import (
	"github.com/gobeetle/fetch/internal/iface"
)

// WithLocation sets the URL and query parameters for the request
func (r *RequestHandler) WithLocation(location iface.RequestLocationHandler) *RequestHandler {
	r.RequestLocationHandler = location
	return r
}

// WithMethod sets the HTTP method for the request
func (r *RequestHandler) WithMethod(method iface.RequestMethodHandler) *RequestHandler {
	r.RequestMethodHandler = method
	return r
}

// WithRequest sets the request
func (r *RequestHandler) WithRequest(request iface.RequestRequestHandler) *RequestHandler {
	r.RequestRequestHandler = request
	return r
}

// WithClient sets a custom HTTP client for the request
func (r *RequestHandler) WithClient(client iface.RequestClientHandler) *RequestHandler {
	r.RequestClientHandler = client
	return r
}

// WithHeader sets the request headers
func (r *RequestHandler) WithHeader(header iface.RequestHeaderHandler) *RequestHandler {
	r.RequestHeaderHandler = header
	return r
}

// WithBody sets the request body
func (r *RequestHandler) WithBody(body iface.RequestBodyHandler) *RequestHandler {
	r.RequestBodyHandler = body
	return r
}
