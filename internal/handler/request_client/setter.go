package request_client

import (
	"net/http"
)

func (r *RequestClient) WithClient(client *http.Client) {
	r.Client = client
}

func (r *RequestClient) Do(request *http.Request) (*http.Response, error) {
	return r.Client.Do(request)
}
