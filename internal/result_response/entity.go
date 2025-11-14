package result_response

import "net/http"

type ResponseResult struct {
	Response   *http.Response
	StatusCode int
	RespBytes  []byte
}
