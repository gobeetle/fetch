package response_err

import errpkg "github.com/gobeetle/fetch/internal/err"

func (r *ResponseErr) GetRetryableFunc() errpkg.GetRetryableFunc {
	return r.RetryableFunc
}
