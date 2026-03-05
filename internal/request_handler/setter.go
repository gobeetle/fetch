package request_handler

import (
	"github.com/gobeetle/fetch/internal/iface"
)

// WithLocation sets the URL and query parameters for the request
func (r *RequestHandler[Req]) WithLocation(location iface.RequestLocationHandler) *RequestHandler[Req] {
	r.RequestLocationHandler = location
	return r
}

// WithMethod sets the HTTP method for the request
func (r *RequestHandler[Req]) WithMethod(method iface.RequestMethodHandler) *RequestHandler[Req] {
	r.RequestMethodHandler = method
	return r
}

// WithRequest sets the request
func (r *RequestHandler[Req]) WithRequest(request iface.RequestRequestHandler) *RequestHandler[Req] {
	r.RequestRequestHandler = request
	return r
}

// WithClient sets a custom HTTP client for the request
func (r *RequestHandler[Req]) WithClient(client iface.RequestClientHandler) *RequestHandler[Req] {
	r.RequestClientHandler = client
	return r
}

// WithHeader sets the request headers
func (r *RequestHandler[Req]) WithHeader(header iface.RequestHeaderHandler) *RequestHandler[Req] {
	r.RequestHeaderHandler = header
	return r
}

// WithBody sets the request body
func (r *RequestHandler[Req]) WithBody(body iface.RequestBodyHandler[Req]) *RequestHandler[Req] {
	r.RequestBodyHandler = body
	return r
}

// WithCurl sets the curl generator handler
func (r *RequestHandler[Req]) WithCurl(curl iface.RequestCurlHandler) *RequestHandler[Req] {
	r.RequestCurlHandler = curl
	return r
}
