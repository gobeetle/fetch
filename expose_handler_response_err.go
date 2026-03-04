package fetch

import "github.com/gobeetle/fetch/internal/handler/response_err"

type (
	ResponseErr = response_err.ResponseErr
)

const ()

var ()

func NewResponseErr() *ResponseErr {
	return response_err.New()
}
