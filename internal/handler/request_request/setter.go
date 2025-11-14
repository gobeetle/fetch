package request_request

import "net/http"

func (r *RequestRequest) WithRequest(request *http.Request) {
	r.Request = request
}
