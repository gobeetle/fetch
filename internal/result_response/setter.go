package result_response

import (
	"net/http"
)

func (r *ResponseResult) SetResponse(response *http.Response) {
	r.response = response
}

func (r *ResponseResult) SetHttpError(httpError error) {
	r.httpError = httpError
}

func (r *ResponseResult) SetStatusCode(statusCode int) {
	r.statusCode = statusCode
}

func (r *ResponseResult) SetRespBytes(respBytes []byte) {
	r.respBytes = respBytes
}

// func (r *ResponseResult) SetRetryable(retryable *retry.Retryable) {
// 	r.retryable = retryable
// }

func (r *ResponseResult) SetAllowMoreRetries(allow bool) {
	r.retryable.SetAllowMoreRetries(allow)
}

func (r *ResponseResult) SetByPassRetryCountCheck(byPass bool) {
	r.retryable.SetByPassRetryCountCheck(byPass)
}
