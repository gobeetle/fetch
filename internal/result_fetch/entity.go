package result_fetch

import (
	rreq "github.com/gobeetle/fetch/internal/result_request"
	rrsq "github.com/gobeetle/fetch/internal/result_response"
)

type Result struct {
	RequestResult  rreq.RequestResult  `json:"request"`
	ResponseResult rrsq.ResponseResult `json:"response"`
}
