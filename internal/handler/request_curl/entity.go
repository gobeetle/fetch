package request_curl

import (
	"github.com/gobeetle/fetch/internal/iface"
)

type RequestCurl struct{}

func (r *RequestCurl) Validate() error {
	return nil
}

var (
	_ iface.RequestCurlHandler = (*RequestCurl)(nil)
)
