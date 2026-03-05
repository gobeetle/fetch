package request_request

import (
	"fmt"
	"net/http"

	errpkg "github.com/gobeetle/fetch/internal/err"
	rreq "github.com/gobeetle/fetch/internal/result_request"
	"github.com/gobeetle/fetch/internal/utils/http2curl"
)

func (r *RequestRequest) PrepareRequest(result *rreq.RequestResult) *errpkg.Error {
	req, err := http.NewRequest(result.GetMethod(), result.GetURL(), nil)
	if err != nil {
		return errpkg.NewError(fmt.Errorf("unable to create request: %w", err))
	}
	result.SetRequest(req)
	command, err := http2curl.GetCurlCommand(req)
	if err != nil {
		result.SetCurl("")
	} else {
		result.SetCurl(command.String())
	}
	return nil
}
