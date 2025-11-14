package response_body

import "github.com/gobeetle/fetch/internal/iface"

type ResponseBody struct {
	Body                 any // response body
	WrapperDataFieldName string
}

var (
	_ iface.ResponseBodyHandler = (*ResponseBody)(nil)
)
