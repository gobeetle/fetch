package response_body

func New[Rsp any]() *ResponseBody[Rsp] {
	return &ResponseBody[Rsp]{}
}
