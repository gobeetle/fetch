package fetch

import "github.com/gobeetle/fetch/internal/handler/response_rawbody"

type (
	ResponseRawBody = response_rawbody.ResponseRawBody
)

const ()

var ()

func NewResponseRawBody() *ResponseRawBody {
	return response_rawbody.New()
}
