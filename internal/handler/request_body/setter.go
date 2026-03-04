package request_body

func (r *RequestBody[Req]) WithJsonBody(body *Req) {
	r.JsonBody = body
}

func (r *RequestBody[Req]) WithRawBody(body []byte) {
	r.RawBody = body
}

func (r *RequestBody[Req]) WithForm(body map[string]any) {
	r.FormBody = body
}
