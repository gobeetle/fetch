package fetch

import "github.com/gobeetle/fetch/internal/handler/response_code"

type (
	ResponseCode = response_code.ResponseCode
)

const ()

var ()

func NewResponseCode() *ResponseCode {
	return response_code.New()
}
