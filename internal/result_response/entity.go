package result_response

import (
	"net/http"

	"github.com/gobeetle/fetch/internal/retry"
)

type ResponseResult struct {
	response   *http.Response
	httpError  error
	statusCode int
	respBytes  []byte
	retryable  retry.Retryable
}
