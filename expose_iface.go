package fetch

import "github.com/gobeetle/fetch/internal/iface"

type (
	// Request handlers
	RequestBodyHandler[Req any] = iface.RequestBodyHandler[Req]
	RequestClientHandler        = iface.RequestClientHandler
	RequestHeaderHandler        = iface.RequestHeaderHandler
	RequestLocationHandler      = iface.RequestLocationHandler
	RequestMethodHandler        = iface.RequestMethodHandler

	// Response handlers
	ResponseBodyHandler[Rsp any] = iface.ResponseBodyHandler[Rsp]
	ResponseCodeHandler          = iface.ResponseCodeHandler
	ResponseErrorHandler         = iface.ResponseErrorHandler

	// Base interfaces
	RequestComponentHandler  = iface.RequestComponentHandler
	ResponseComponentHandler = iface.ResponseComponentHandler
	Validator                = iface.Validator
)

const ()

var ()
