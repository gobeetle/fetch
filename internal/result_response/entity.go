package result_response

import (
	"net/http"

	"github.com/gobeetle/fetch/internal/retry"
)

type ResponseResult struct {
	Response   *http.Response
	HttpError  error
	StatusCode int
	RespBytes  []byte
	*retry.Retryable
}
