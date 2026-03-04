package result_response

import (
	"net/http"

	"github.com/gobeetle/fetch/internal/retry"
)

func New() ResponseResult {
	r := ResponseResult{
		response:   &http.Response{},
		statusCode: 0,
		respBytes:  []byte{},
		retryable:  *retry.NewRetryable(),
	}
	return r
}
