package response_body

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

func (r *ResponseBody[Rsp]) ValidateResponse(result *rrsp.ResponseResult) *errpkg.Error {
	return nil
}

func (r *ResponseBody[Rsp]) Validate() error {
	return nil
}
