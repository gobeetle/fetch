package request_method

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	rreq "github.com/gobeetle/fetch/internal/result_request"
)

func (r *RequestMethod) PrepareRequest(result *rreq.RequestResult) *errpkg.Error {
	result.SetMethod(string(r.Method))
	return nil
}
