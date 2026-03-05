// Package defaults provides default handler creators and allows customization of default handlers.
package defaults

import (
	"reflect"
	"sync"

	"github.com/gobeetle/fetch/internal/handler/request_body"
	"github.com/gobeetle/fetch/internal/handler/request_client"
	"github.com/gobeetle/fetch/internal/handler/request_curl"
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
	mu                      sync.RWMutex
	defaultRequestCreators  = map[reflect.Type]any{}
	defaultResponseCreators = map[reflect.Type]any{}
)

// RequestHandlerCreator defines the function signature for creating a new request handler
type RequestHandlerCreator[Req any] func() *hreq.RequestHandler[Req]

// ResponseHandlerCreator defines the function signature for creating a new response handler
type ResponseHandlerCreator[Rsp any] func() *hrsp.ResponseHandler[Rsp]

func typeKey[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

// NewRequestHandler returns a new request handler with default configuration
func NewRequestHandler[Req any]() *hreq.RequestHandler[Req] {
	mu.RLock()
	defer mu.RUnlock()
	if v, ok := defaultRequestCreators[typeKey[Req]()]; ok {
		if creator, ok := v.(RequestHandlerCreator[Req]); ok {
			return creator()
		}
	}
	return defaultRequestHandler[Req]()
}

// NewResponseHandler returns a new response handler with default configuration
func NewResponseHandler[Rsp any]() *hrsp.ResponseHandler[Rsp] {
	mu.RLock()
	defer mu.RUnlock()
	if v, ok := defaultResponseCreators[typeKey[Rsp]()]; ok {
		if creator, ok := v.(ResponseHandlerCreator[Rsp]); ok {
			return creator()
		}
	}
	return defaultResponseHandler[Rsp]()
}

// SetDefaultRequestHandler sets the default request handler creator
// Pass nil to reset to the default implementation
func SetDefaultRequestHandler[Req any](creator RequestHandlerCreator[Req]) {
	mu.Lock()
	defer mu.Unlock()
	key := typeKey[Req]()
	if creator == nil {
		delete(defaultRequestCreators, key)
		return
	}
	defaultRequestCreators[key] = creator
}

// SetDefaultResponseHandler sets the default response handler creator
// Pass nil to reset to the default implementation
func SetDefaultResponseHandler[Rsp any](creator ResponseHandlerCreator[Rsp]) {
	mu.Lock()
	defer mu.Unlock()
	key := typeKey[Rsp]()
	if creator == nil {
		delete(defaultResponseCreators, key)
		return
	}
	defaultResponseCreators[key] = creator
}

// defaultRequestHandler is the default implementation of request handler creator
func defaultRequestHandler[Req any]() *hreq.RequestHandler[Req] {
	return new(hreq.RequestHandler[Req]).
		WithLocation(request_location.New()).
		WithMethod(request_method.New()).
		WithRequest(request_request.New()).
		WithClient(request_client.New()).
		WithHeader(request_header.New()).
		WithBody(request_body.New[Req]()).
		WithCurl(request_curl.New())
}

// defaultResponseHandler is the default implementation of response handler creator
func defaultResponseHandler[Rsp any]() *hrsp.ResponseHandler[Rsp] {
	return new(hrsp.ResponseHandler[Rsp]).
		WithErrorHandler(response_err.New()).
		WithCodeHandler(response_code.New()).
		WithRawBodyHandler(response_rawbody.New()).
		WithBodyHandler(response_body.New[Rsp]())
}
