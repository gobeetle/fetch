package request_curl

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	rreq "github.com/gobeetle/fetch/internal/result_request"
	"github.com/gobeetle/fetch/internal/utils/http2curl"
)

func (r *RequestCurl) PrepareRequest(result *rreq.RequestResult) *errpkg.Error {
	req := result.GetRequest()
	if req == nil {
		result.SetCurl("")
		return nil
	}
	command, err := http2curl.GetCurlCommand(req)
	if err != nil {
		result.SetCurl("")
		return nil
	}
	result.SetCurl(command.String())
	return nil
}
