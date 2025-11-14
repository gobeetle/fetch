// Package defaults provides default handler creators and allows customization of default handlers.
package defaults

import (
	"sync"

	"github.com/gobeetle/fetch/internal/handler/request_body"
	"github.com/gobeetle/fetch/internal/handler/request_client"
	"github.com/gobeetle/fetch/internal/handler/request_header"
	"github.com/gobeetle/fetch/internal/handler/request_location"
	"github.com/gobeetle/fetch/internal/handler/request_method"
	"github.com/gobeetle/fetch/internal/handler/request_request"
	"github.com/gobeetle/fetch/internal/handler/response_body"
	"github.com/gobeetle/fetch/internal/handler/response_code"
	"github.com/gobeetle/fetch/internal/handler/response_err"
	"github.com/gobeetle/fetch/internal/handler/response_rawbody"
	hreq "github.com/gobeetle/fetch/internal/request_handler"
	hrsp "github.com/gobeetle/fetch/internal/response_handler"
)

var (
	mu                     sync.RWMutex
	defaultRequestCreator  = defaultRequestHandler
	defaultResponseCreator = defaultResponseHandler
)

// RequestHandlerCreator defines the function signature for creating a new request handler
type RequestHandlerCreator func() *hreq.RequestHandler

// ResponseHandlerCreator defines the function signature for creating a new response handler
type ResponseHandlerCreator func() *hrsp.ResponseHandler

// NewRequestHandler returns a new request handler with default configuration
func NewRequestHandler() *hreq.RequestHandler {
	mu.RLock()
	defer mu.RUnlock()
	return defaultRequestCreator()
}

// NewResponseHandler returns a new response handler with default configuration
func NewResponseHandler() *hrsp.ResponseHandler {
	mu.RLock()
	defer mu.RUnlock()
	return defaultResponseCreator()
}

// SetDefaultRequestHandler sets the default request handler creator
// Pass nil to reset to the default implementation
func SetDefaultRequestHandler(creator RequestHandlerCreator) {
	mu.Lock()
	defer mu.Unlock()
	if creator == nil {
		defaultRequestCreator = defaultRequestHandler
	} else {
		defaultRequestCreator = creator
	}
}

// SetDefaultResponseHandler sets the default response handler creator
// Pass nil to reset to the default implementation
func SetDefaultResponseHandler(creator ResponseHandlerCreator) {
	mu.Lock()
	defer mu.Unlock()
	if creator == nil {
		defaultResponseCreator = defaultResponseHandler
	} else {
		defaultResponseCreator = creator
	}
}

// defaultRequestHandler is the default implementation of request handler creator
func defaultRequestHandler() *hreq.RequestHandler {
	return new(hreq.RequestHandler).
		WithLocation(request_location.New()).
		WithMethod(request_method.New()).
		WithRequest(request_request.New()).
		WithClient(request_client.New()).
		WithHeader(request_header.New()).
		WithBody(request_body.New())
}

// defaultResponseHandler is the default implementation of response handler creator
func defaultResponseHandler() *hrsp.ResponseHandler {
	return new(hrsp.ResponseHandler).
		WithErrorHandler(response_err.New()).
		WithCodeHandler(response_code.New()).
		WithRawBodyHandler(response_rawbody.New()).
		WithBodyHandler(response_body.New())
}
