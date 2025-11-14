package fetch

import "github.com/gobeetle/fetch/internal/handler/request_client"

type (
	RequestClient = request_client.RequestClient
)

const ()

var (
	NewRequestClient = request_client.New
)
