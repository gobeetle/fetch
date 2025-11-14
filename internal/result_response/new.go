package result_response

import "net/http"

func New() ResponseResult {
	return ResponseResult{
		Response:   &http.Response{},
		StatusCode: 0,
		RespBytes:  []byte{},
	}
}
