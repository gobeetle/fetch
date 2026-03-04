package fetch

import (
	"net/http"
	"time"

	errpkg "github.com/gobeetle/fetch/internal/err"
	hreq "github.com/gobeetle/fetch/internal/request_handler"
	mreq "github.com/gobeetle/fetch/internal/request_modifier"
	hrsp "github.com/gobeetle/fetch/internal/response_handler"
	mrsp "github.com/gobeetle/fetch/internal/response_modifier"
	rfetch "github.com/gobeetle/fetch/internal/result_fetch"
	rreq "github.com/gobeetle/fetch/internal/result_request"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
	"github.com/gobeetle/fetch/internal/retry"
)

type Fetch[Req, Rsp any] struct {
	req_handler *hreq.RequestHandler[Req]
	rsp_handler *hrsp.ResponseHandler[Rsp]
	retry       *retry.Retry[*rreq.RequestResult, rrsp.ResponseResult]
}

func (f *Fetch[Req, Rsp]) WithRetry(count int, delay time.Duration) *Fetch[Req, Rsp] {
	f.retry.WithRetry(count, delay)
	return f
}

/*
ModReq is a function that modifies the request handler
functions are predefined in the fetch package
modifiers are applied in the order they are passed
*/
func (f *Fetch[Req, Rsp]) ModReq(modifiers ...mreq.RequestModifier[Req]) *Fetch[Req, Rsp] {
	f.ensure_request_handler()
	mreq.Modify(f.req_handler, modifiers...)
	return f
}

/*
ModRsp is a function that modifies the response handler
functions are predefined in the fetch package
modifiers are applied in the order they are passed
*/
func (f *Fetch[Req, Rsp]) ModRsp(modifiers ...mrsp.ResponseModifier[Rsp]) *Fetch[Req, Rsp] {
	f.ensure_response_handler()
	mrsp.Modify(f.rsp_handler, modifiers...)
	return f
}

// Do performs the HTTP request with retry logic
func (f *Fetch[Req, Rsp]) Do() (rfetch.Result, *errpkg.Error) {
	f.ensure_retry()
	f.ensure_request_handler()
	f.ensure_response_handler()
	fetch_r := rfetch.New()

	// prepare the request
	if r, err := f.req_handler.Request(); err != nil {
		fetch_r.RequestResult = r
		return fetch_r, err
	} else {
		fetch_r.RequestResult = r
	}

	// execute the request with retry logic
	if r, err := f.retry.Do(f.do, &fetch_r.RequestResult); err != nil {
		fetch_r.ResponseResult = r
		return fetch_r, errpkg.NewError(err)
	} else {
		fetch_r.ResponseResult = r
	}

	return fetch_r, nil
}

// do the request
func (f *Fetch[Req, Rsp]) do(request_result *rreq.RequestResult) (rrsp.ResponseResult, error) {
	var response *http.Response
	rsp_r := rrsp.New()
	// do request
	response, err := f.req_handler.RequestClientHandler.Do(request_result.GetRequest())
	rsp_r.SetResponse(response)
	rsp_r.SetHttpError(err)
	// we will not fail the request if the http call fails, we will let the response handler deal with it
	// prepare the response
	if err := f.rsp_handler.Response(&rsp_r); err != nil {
		return rsp_r, err
	} else {
		return rsp_r, nil
	}
}
