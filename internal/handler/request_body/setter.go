package request_body

func (r *RequestBody) WithJsonBody(body any) {
	r.JsonBody = body
}

func (r *RequestBody) WithRawBody(body []byte) {
	r.RawBody = body
}

func (r *RequestBody) WithForm(body map[string]any) {
	r.FormBody = body
}
