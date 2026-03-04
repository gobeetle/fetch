package fetch

import (
	rresp "github.com/gobeetle/fetch/internal/result_response"
)

type (
	ResponseResult = rresp.ResponseResult
)

const ()

var ()

func NewResponseResult() ResponseResult {
	return rresp.New()
}
