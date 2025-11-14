package request_body

import "github.com/gobeetle/fetch/internal/iface"

type RequestBody struct {
	JsonBody any
	RawBody  []byte
	FormBody map[string]any
}

var (
	_ iface.RequestBodyHandler = (*RequestBody)(nil)
)
