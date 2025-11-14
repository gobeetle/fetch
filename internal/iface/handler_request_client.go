package iface

import "net/http"

type RequestClientHandler interface {
	RequestComponentHandler
	WithClient(client *http.Client)
	Do(request *http.Request) (*http.Response, error)
}
