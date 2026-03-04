package result_fetch

import (
	rreq "github.com/gobeetle/fetch/internal/result_request"
	rrsq "github.com/gobeetle/fetch/internal/result_response"
)

func New() Result {
	return Result{
		RequestResult:  rreq.New(),
		ResponseResult: rrsq.New(),
	}
}
