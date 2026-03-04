package request_modifier

import (
	hreq "github.com/gobeetle/fetch/internal/request_handler"
)

/*
RequestModifier is a function that modifies the request handler
functions are predefined in the package
*/
type RequestModifier[Req any] func(*hreq.RequestHandler[Req]) *hreq.RequestHandler[Req]

func Modify[Req any](handler *hreq.RequestHandler[Req], modifiers ...RequestModifier[Req]) *hreq.RequestHandler[Req] {
	for _, m := range modifiers {
		if m == nil {
			continue
		}
		m(handler)
	}
	return handler
}
