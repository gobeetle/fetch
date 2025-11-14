package request_modifier

import (
	hreq "github.com/gobeetle/fetch/internal/request_handler"
)

/*
RequestModifier is a function that modifies the request handler
functions are predefined in the package
*/
type RequestModifier func(*hreq.RequestHandler) *hreq.RequestHandler

func Modify(handler *hreq.RequestHandler, modifiers ...RequestModifier) *hreq.RequestHandler {
	for _, m := range modifiers {
		if m == nil {
			continue
		}
		m(handler)
	}
	return handler
}
