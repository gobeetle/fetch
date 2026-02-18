package response_err

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	"github.com/gobeetle/fetch/internal/iface"
)

type ResponseErr struct {
	RetryableFuncs    []errpkg.GetRetryableFunc
	UseErrorUnwrapper bool
	Err               error
}

func (r *ResponseErr) Error() error {
	return r.Err
}

var (
	_ iface.ResponseErrorHandler = (*ResponseErr)(nil)
)
