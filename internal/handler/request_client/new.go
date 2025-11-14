package request_client

import "net/http"

func New() *RequestClient {
	return &RequestClient{
		Client: &http.Client{},
	}
}
