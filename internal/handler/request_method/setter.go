package request_method

import "github.com/gobeetle/fetch/internal/enum"

func (r *RequestMethod) WithMethod(method enum.MethodType) {
	r.Method = method
}

func (r *RequestMethod) WithMethodString(method string) {
	r.Method = enum.MethodType(method)
}
