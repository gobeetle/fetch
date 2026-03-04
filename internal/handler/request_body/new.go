package request_body

func New[Req any]() *RequestBody[Req] {
	return &RequestBody[Req]{}
}
