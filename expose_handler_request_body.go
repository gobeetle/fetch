package fetch

import "github.com/gobeetle/fetch/internal/handler/request_body"

type (
	RequestBody[Req any] = request_body.RequestBody[Req]
)

const ()

var ()

func NewRequestBody[Req any]() *RequestBody[Req] {
	return request_body.New[Req]()
}
