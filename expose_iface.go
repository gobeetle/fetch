package fetch

import "github.com/gobeetle/fetch/internal/iface"

type (
	// Request handlers
	RequestBodyHandler     = iface.RequestBodyHandler
	RequestClientHandler   = iface.RequestClientHandler
	RequestHeaderHandler   = iface.RequestHeaderHandler
	RequestLocationHandler = iface.RequestLocationHandler
	RequestMethodHandler   = iface.RequestMethodHandler

	// Response handlers
	ResponseBodyHandler   = iface.ResponseBodyHandler
	ResponseCodeHandler   = iface.ResponseCodeHandler
	ResponseErrorHandler  = iface.ResponseErrorHandler

	// Base interfaces
	RequestComponentHandler  = iface.RequestComponentHandler
	ResponseComponentHandler = iface.ResponseComponentHandler
	Validator               = iface.Validator
)

const ()

var ()
