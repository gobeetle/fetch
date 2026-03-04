package fetch

import "github.com/gobeetle/fetch/internal/handler/request_method"

type (
	RequestMethod = request_method.RequestMethod
)

const ()

var ()

func NewRequestMethod() *RequestMethod {
	return request_method.New()
}
