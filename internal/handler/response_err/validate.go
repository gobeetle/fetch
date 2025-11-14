package response_err

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

func (r *ResponseErr) ValidateResponse(result *rrsp.ResponseResult) *errpkg.Error {
	return nil
}

func (r *ResponseErr) Validate() error {
	return nil
}
