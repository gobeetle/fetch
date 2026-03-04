package fetch

import (
	rreq "github.com/gobeetle/fetch/internal/result_request"
)

type (
	RequestResult = rreq.RequestResult
)

const ()

var ()

func NewRequestResult() RequestResult {
	return rreq.New()
}
