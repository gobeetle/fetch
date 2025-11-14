package iface

import "net/http"

type RequestRequestHandler interface {
	RequestComponentHandler
	WithRequest(request *http.Request)
}
