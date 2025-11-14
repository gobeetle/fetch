package result_request

import "net/http"

type RequestResult struct {
	Method  string
	URL     string
	Request *http.Request
}
