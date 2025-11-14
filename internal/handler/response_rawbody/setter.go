package response_rawbody

func (r *ResponseRawBody) WithRawBodyBytes(body_bytes []byte) {
	r.BodyBytes = body_bytes
}
