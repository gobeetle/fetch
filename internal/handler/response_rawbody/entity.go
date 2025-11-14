package response_rawbody

import "github.com/gobeetle/fetch/internal/iface"

type ResponseRawBody struct {
	BodyBytes  []byte
}

var (
	_ iface.ResponseRawBodyHandler = (*ResponseRawBody)(nil)
)
