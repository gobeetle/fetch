package fetch

import (
	"github.com/gobeetle/fetch/internal/handler/defaults"
	hreq "github.com/gobeetle/fetch/internal/request_handler"
	hrsp "github.com/gobeetle/fetch/internal/response_handler"
	rreq "github.com/gobeetle/fetch/internal/result_request"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
	"github.com/gobeetle/fetch/internal/retry"
)

// New creates a new Fetch instance with default handlers
func New() *Fetch {
	r := &Fetch{}
	r.ensure_request_handler()
	r.ensure_response_handler()
	r.ensure_retry()
	return r
}

func (f *Fetch) ensure_request_handler() *hreq.RequestHandler {
	if f.req_handler == nil {
		f.req_handler = defaults.NewRequestHandler()
	}
	return f.req_handler
}

func (f *Fetch) ensure_response_handler() *hrsp.ResponseHandler {
	if f.rsp_handler == nil {
		f.rsp_handler = defaults.NewResponseHandler()
	}
	return f.rsp_handler
}

func (f *Fetch) ensure_retry() *retry.Retry[*rreq.RequestResult, rrsp.ResponseResult] {
	if f.retry == nil {
		f.retry = retry.New[*rreq.RequestResult, rrsp.ResponseResult]()
	}
	return f.retry
}
