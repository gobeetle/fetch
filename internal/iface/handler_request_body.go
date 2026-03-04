package iface

type RequestBodyHandler[Req any] interface {
	RequestComponentHandler
	WithJsonBody(body *Req)
	WithRawBody(body []byte)
	WithForm(body map[string]any)
}
