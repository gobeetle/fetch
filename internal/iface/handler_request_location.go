package iface

type RequestLocationHandler interface {
	RequestComponentHandler
	WithUrl(endpoint string)
	WithQueryParams(params map[string]string)
	WithQueryParamPair(key string, value string)
	WithQueryParamsCleared()
}
