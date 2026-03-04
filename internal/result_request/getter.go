package result_request

import (
	"net/http"

	"github.com/gobeetle/fetch/internal/utils/http2curl"
)

func (r *RequestResult) GetMethod() string {
	return r.method
}

func (r *RequestResult) GetURL() string {
	return r.url
}

func (r *RequestResult) GetRequest() *http.Request {
	return r.request
}

func (r *RequestResult) Curl() string {
	if r.request == nil {
		return ""
	}
	command, err := http2curl.GetCurlCommand(r.request)
	if err != nil {
		return ""
	}
	return command.String()
}
