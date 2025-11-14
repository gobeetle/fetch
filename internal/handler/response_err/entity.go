package response_err

import "github.com/gobeetle/fetch/internal/iface"

type ResponseErr struct {
	UseErrorUnwrapper bool
	Err               error
}

func (r *ResponseErr) Error() error {
	return r.Err
}

var (
	_ iface.ResponseErrorHandler = (*ResponseErr)(nil)
)
