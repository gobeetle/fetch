package fetch

import "github.com/gobeetle/fetch/internal/handler/request_client"

type (
	RequestClient = request_client.RequestClient
)

const ()

var ()

func NewRequestClient() *RequestClient {
	return request_client.New()
}
