package fetch

import "github.com/gobeetle/fetch/internal/handler/defaults"

type (
	RequestHandlerCreator  = defaults.RequestHandlerCreator
	ResponseHandlerCreator = defaults.ResponseHandlerCreator
)

const ()

var (
	NewRequestHandler         = defaults.NewRequestHandler
	NewResponseHandler        = defaults.NewResponseHandler
	SetDefaultRequestHandler  = defaults.SetDefaultRequestHandler
	SetDefaultResponseHandler = defaults.SetDefaultResponseHandler
)
