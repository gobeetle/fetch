package response_rawbody

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

func (r *ResponseRawBody) ValidateResponse(result *rrsp.ResponseResult) *errpkg.Error {
	return nil
}

func (r *ResponseRawBody) Validate() error {
	return nil
}
