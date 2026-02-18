package result_response

import (
	"net/http"

	"github.com/gobeetle/fetch/internal/retry"
)

func New() ResponseResult {
	r := ResponseResult{
		Response:   &http.Response{},
		StatusCode: 0,
		RespBytes:  []byte{},
		Retryable:  retry.NewRetryable(),
	}
	return r
}
