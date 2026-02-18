package response_err

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
)

func (r *ResponseErr) WithError(err error) {
	r.Err = err
}

func (r *ResponseErr) WithUseErrorUnwrapper() {
	r.UseErrorUnwrapper = true
}

func (r *ResponseErr) WithRetryableFunc(fns ...errpkg.GetRetryableFunc) {
	if len(fns) == 0 {
		return
	}
	r.RetryableFuncs = append(r.RetryableFuncs, fns...)
}
