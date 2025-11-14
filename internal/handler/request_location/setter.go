package request_location

func (r *RequestLocation) WithUrl(endpoint string) {
	r.URL = endpoint
}

func (r *RequestLocation) WithQueryParams(params map[string]string) {
	r.QueryParams = params
}

func (r *RequestLocation) WithQueryParamPair(key string, value string) {
	if r.QueryParams == nil {
		r.QueryParams = make(map[string]string)
	}
	r.QueryParams[key] = value
}

func (r *RequestLocation) WithQueryParamsCleared() {
	r.QueryParams = make(map[string]string)
}
