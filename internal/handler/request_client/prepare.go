package request_client

import (
	"net/http"

	errpkg "github.com/gobeetle/fetch/internal/err"
	rreq "github.com/gobeetle/fetch/internal/result_request"
)

func (r *RequestClient) PrepareRequest(_ *rreq.RequestResult) *errpkg.Error {
	if r.Client == nil {
		r.Client = &http.Client{}
	}
	return nil
}
