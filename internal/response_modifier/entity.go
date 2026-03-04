package response_modifier

import hrsp "github.com/gobeetle/fetch/internal/response_handler"

/*
rsp_opt_modifier is a function that modifies the response handler
functions are predefined in the fetch package
*/
type ResponseModifier[Rsp any] func(*hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp]

func Modify[Rsp any](handler *hrsp.ResponseHandler[Rsp], modifiers ...ResponseModifier[Rsp]) *hrsp.ResponseHandler[Rsp] {
	for _, m := range modifiers {
		if m == nil {
			continue
		}
		m(handler)
	}
	return handler
}
