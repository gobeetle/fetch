package response_body

func (r *ResponseBody[Rsp]) WithJsonObj(resp_obj *Rsp) {
	r.Body = resp_obj
}
