package fetch

import (
	"net/http"
	"time"

	errpkg "github.com/gobeetle/fetch/internal/err"
	hreq "github.com/gobeetle/fetch/internal/request_handler"
	mreq "github.com/gobeetle/fetch/internal/request_modifier"
	hrsp "github.com/gobeetle/fetch/internal/response_handler"
	mrsp "github.com/gobeetle/fetch/internal/response_modifier"
	rreq "github.com/gobeetle/fetch/internal/result_request"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
	"github.com/gobeetle/fetch/internal/retry"
)

type Fetch struct {
	req_handler *hreq.RequestHandler
	rsp_handler *hrsp.ResponseHandler
	retry       *retry.Retry[*rreq.RequestResult, rrsp.ResponseResult]
}

func (f *Fetch) WithRetry(count int, delay time.Duration) *Fetch {
	f.retry.WithRetry(count, delay)
	return f
}

/*
ModReq is a function that modifies the request handler
functions are predefined in the fetch package
modifiers are applied in the order they are passed
*/
func (f *Fetch) ModReq(modifiers ...mreq.RequestModifier) *Fetch {
	f.ensure_request_handler()
	mreq.Modify(f.req_handler, modifiers...)
	return f
}

/*
ModRsp is a function that modifies the response handler
functions are predefined in the fetch package
modifiers are applied in the order they are passed
*/
func (f *Fetch) ModRsp(modifiers ...mrsp.ResponseModifier) *Fetch {
	f.ensure_response_handler()
	mrsp.Modify(f.rsp_handler, modifiers...)
	return f
}

// Do performs the HTTP request with retry logic
func (f *Fetch) Do() (rrsp.ResponseResult, *errpkg.Error) {
	f.ensure_retry()
	f.ensure_request_handler()
	f.ensure_response_handler()
	req_r := rreq.New()
	rsp_r := rrsp.New()

	// prepare the request
	if r, err := f.req_handler.Request(); err != nil {
		return rsp_r, err
	} else {
		req_r = r
	}

	// execute the request with retry logic
	if r, err := f.retry.Do(f.do, &req_r); err != nil {
		return rsp_r, errpkg.NewError(err)
	} else {
		rsp_r = r
	}

	return rsp_r, nil
}

// do the request
func (f *Fetch) do(request_result *rreq.RequestResult) (rrsp.ResponseResult, error) {
	var response *http.Response
	rsp_r := rrsp.New()
	// do request
	response, err := f.req_handler.RequestClientHandler.Do(request_result.Request)
	rsp_r.Response = response
	rsp_r.HttpError = err
	// we will not fail the request if the http call fails, we will let the response handler deal with it
	// prepare the response
	if err := f.rsp_handler.Response(&rsp_r); err != nil {
		return rsp_r, err
	} else {
		return rsp_r, nil
	}
}
