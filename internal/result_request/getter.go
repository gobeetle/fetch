package result_request

import (
	"net/http"
)

func (r *RequestResult) IsEmpty() bool {
	return r.method == "" && r.url == "" && r.curl == "" && r.request == nil
}

func (r *RequestResult) GetMethod() string {
	return r.method
}

func (r *RequestResult) GetURL() string {
	return r.url
}

func (r *RequestResult) GetRequest() *http.Request {
	return r.request
}

func (r *RequestResult) GetCurl() string {
	return r.curl
}
