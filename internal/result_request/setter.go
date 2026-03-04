package result_request

import "net/http"

func (r *RequestResult) SetMethod(method string) {
	r.method = method
}

func (r *RequestResult) SetURL(url string) {
	r.url = url
}

func (r *RequestResult) SetRequest(request *http.Request) {
	r.request = request
}
