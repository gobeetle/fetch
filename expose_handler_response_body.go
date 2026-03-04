package fetch

import "github.com/gobeetle/fetch/internal/handler/response_body"

type (
	ResponseBody[Rsp any] = response_body.ResponseBody[Rsp]
)

const ()

var ()

func NewResponseBody[Rsp any]() *ResponseBody[Rsp] {
	return response_body.New[Rsp]()
}
