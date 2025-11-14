package response_code

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

func (r *ResponseCode) Prepare(result *rrsp.ResponseResult) {
	// if status code is set, use it instead of reading from response
	if r.StatusCode != nil {
		result.StatusCode = *r.StatusCode
	}
	if rsp := result.Response; rsp != nil {
		result.StatusCode = rsp.StatusCode
	}
}

func (r *ResponseCode) PrepareResponse(_ *rrsp.ResponseResult) *errpkg.Error {
	return nil
}
