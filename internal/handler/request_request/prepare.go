package request_request

import (
	"fmt"
	"net/http"

	errpkg "github.com/gobeetle/fetch/internal/err"
	rreq "github.com/gobeetle/fetch/internal/result_request"
)

func (r *RequestRequest) PrepareRequest(result *rreq.RequestResult) *errpkg.Error {
	if r, err := http.NewRequest(result.Method, result.URL, nil); err != nil {
		return errpkg.NewError(fmt.Errorf("unable to create request: %w", err))
	} else {
		result.Request = r
	}
	return nil
}
