package response_body

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

func (r *ResponseBody) ValidateResponse(result *rrsp.ResponseResult) *errpkg.Error {
	return nil
}

func (r *ResponseBody) Validate() error {
	return nil
}
