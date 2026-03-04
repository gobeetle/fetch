package request_request

import (
	"fmt"
	"net/http"

	errpkg "github.com/gobeetle/fetch/internal/err"
	rreq "github.com/gobeetle/fetch/internal/result_request"
)

func (r *RequestRequest) PrepareRequest(result *rreq.RequestResult) *errpkg.Error {
	if req, err := http.NewRequest(result.GetMethod(), result.GetURL(), nil); err != nil {
		return errpkg.NewError(fmt.Errorf("unable to create request: %w", err))
	} else {
		result.SetRequest(req)
	}
	return nil
}
