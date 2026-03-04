package result_request

import "net/http"

func New() RequestResult {
	return RequestResult{
		request: &http.Request{},
	}
}
