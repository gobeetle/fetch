package request_client

import (
	"net/http"

	"github.com/gobeetle/fetch/internal/iface"
)

type RequestClient struct {
	Client *http.Client // http client
}

var (
	_ iface.RequestClientHandler = (*RequestClient)(nil)
)
