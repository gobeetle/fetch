package request_request

import (
	"net/http"

	"github.com/gobeetle/fetch/internal/iface"
)

type RequestRequest struct {
	Request *http.Request
}

var (
	_ iface.RequestRequestHandler = (*RequestRequest)(nil)
)
