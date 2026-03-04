package fetch

import "github.com/gobeetle/fetch/internal/handler/request_header"

type (
	RequestHeader = request_header.RequestHeader
)

const ()

var ()

func NewRequestHeader() *RequestHeader {
	return request_header.New()
}
