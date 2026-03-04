package response_body

import "github.com/gobeetle/fetch/internal/iface"

type ResponseBody[Rsp any] struct {
	Body *Rsp // response body
}

var (
	_ iface.ResponseBodyHandler[*any] = (*ResponseBody[*any])(nil)
)
