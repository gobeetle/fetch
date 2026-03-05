package result_request

import (
	"net/http"
)

type RequestResult struct {
	method  string
	url     string
	curl    string
	request *http.Request
}
