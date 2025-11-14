package fetch

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	hreq "github.com/gobeetle/fetch/internal/request_handler"
)

type (
	RequestHandler = hreq.RequestHandler
)

const ()

var (
	// this helps to build a http request with the use of modifiers
	// it is helpful when you want to build a request with modifiers
	// it is helpful for external integration without using fetch instance
	BuildRequest = func(modifiers ...RequestModifier) (RequestResult, *errpkg.Error) {
		return ModifyRequest(NewRequestHandler(), modifiers...).Request()
	}
)
