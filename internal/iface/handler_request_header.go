package iface

import "github.com/gobeetle/fetch/internal/enum"

type RequestHeaderHandler interface {
	RequestComponentHandler
	WithHeaderContentTypeAsJson()
	WithHeaderContentType(contentType enum.MediaType)
	WithHeaderAuth(token string)
	WithHeader(header map[string]string)
	WithHeaderPair(key string, value string)
	WithHeaderCleared()
}
