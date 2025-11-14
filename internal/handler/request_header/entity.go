package request_header

import (
	"github.com/gobeetle/fetch/internal/enum"
	"github.com/gobeetle/fetch/internal/iface"
)

type RequestHeader struct {
	HeaderContentType enum.MediaType
	Header            map[string]string // request headers
	HeaderAuth        string            // Authorization header
}

var (
	_ iface.RequestHeaderHandler = (*RequestHeader)(nil)
)
