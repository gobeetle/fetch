package request_location

import "github.com/gobeetle/fetch/internal/iface"

type RequestLocation struct {
	URL         string            // URL
	QueryParams map[string]string // query parameters
}

var (
	_ iface.RequestLocationHandler = (*RequestLocation)(nil)
)
