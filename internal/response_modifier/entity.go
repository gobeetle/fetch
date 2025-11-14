package response_modifier

import hrsp "github.com/gobeetle/fetch/internal/response_handler"

/*
rsp_opt_modifier is a function that modifies the response handler
functions are predefined in the fetch package
*/
type ResponseModifier func(*hrsp.ResponseHandler) *hrsp.ResponseHandler

func Modify(handler *hrsp.ResponseHandler, modifiers ...ResponseModifier) *hrsp.ResponseHandler {
	for _, m := range modifiers {
		if m == nil {
			continue
		}
		m(handler)
	}
	return handler
}
