package request_header

import (
	"net/http"

	errpkg "github.com/gobeetle/fetch/internal/err"
	rreq "github.com/gobeetle/fetch/internal/result_request"
)

func (r *RequestHeader) PrepareRequest(result *rreq.RequestResult) *errpkg.Error {
	req := result.Request
	if req.Header == nil {
		req.Header = make(http.Header)
	}
	if r.Header != nil {
		for key, value := range r.Header {
			req.Header.Set(key, value)
		}
	}
	if r.HeaderAuth != "" {
		if req.Header == nil {
			req.Header = make(http.Header)
		}
		req.Header.Set("Authorization", r.HeaderAuth)
	}
	return nil
}
