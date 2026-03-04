package fetch

import (
	"github.com/gobeetle/fetch/internal/handler/defaults"
)

type (
	RequestHandlerCreator[Req any]  = defaults.RequestHandlerCreator[Req]
	ResponseHandlerCreator[Rsp any] = defaults.ResponseHandlerCreator[Rsp]
)

const ()

var ()

func NewRequestHandler[Req any]() *RequestHandler[Req] {
	return defaults.NewRequestHandler[Req]()
}

func SetDefaultRequestHandler[Req any](creator defaults.RequestHandlerCreator[Req]) {
	defaults.SetDefaultRequestHandler(creator)
}

func NewResponseHandler[Rsp any]() *ResponseHandler[Rsp] {
	return defaults.NewResponseHandler[Rsp]()
}

func SetDefaultResponseHandler[Rsp any](creator defaults.ResponseHandlerCreator[Rsp]) {
	defaults.SetDefaultResponseHandler[Rsp](creator)
}
