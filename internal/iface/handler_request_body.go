package iface

type RequestBodyHandler interface {
	RequestComponentHandler
	WithJsonBody(body any)
	WithRawBody(body []byte)
	WithForm(body map[string]any)
}
