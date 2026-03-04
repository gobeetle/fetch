package request_body

import "github.com/gobeetle/fetch/internal/iface"

type RequestBody[Req any] struct {
	JsonBody *Req
	RawBody  []byte
	FormBody map[string]any
}

var (
	_ iface.RequestBodyHandler[any] = (*RequestBody[any])(nil)
)
