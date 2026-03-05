package result_response

import (
	"net/http"
)

func (r *ResponseResult) IsEmpty() bool {
	return r.response == nil
}

func (r *ResponseResult) GetResponse() *http.Response {
	return r.response
}

func (r *ResponseResult) GetHttpError() error {
	return r.httpError
}

func (r *ResponseResult) GetStatusCode() int {
	return r.statusCode
}

func (r *ResponseResult) GetRespBytes() []byte {
	return r.respBytes
}

// func (r *ResponseResult) GetRetryable() *retry.Retryable {
// 	return r.retryable
// }

func (r ResponseResult) AllowMoreRetries() bool {
	return r.retryable.AllowMoreRetries()
}

func (r ResponseResult) ByPassRetryCountCheck() bool {
	return r.retryable.ByPassRetryCountCheck()
}
