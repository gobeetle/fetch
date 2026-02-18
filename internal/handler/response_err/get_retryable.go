package response_err

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
	"github.com/gobeetle/fetch/internal/retry"
)

func (r *ResponseErr) GetRetryableFunc() errpkg.GetRetryableFunc {
	if len(r.RetryableFuncs) == 0 {
		return nil
	}
	return func(result rrsp.ResponseResult, e errpkg.Error) *retry.Retryable {
		allow := true
		bypass := false
		any := false
		for _, fn := range r.RetryableFuncs {
			if fn == nil {
				continue
			}
			rr := fn(result, e)
			if rr == nil {
				continue
			}
			any = true
			allow = allow && rr.AllowMoreRetries()
			bypass = bypass || rr.ByPassRetryCountCheck()
		}
		if !any {
			return nil
		}
		return retry.NewRetryable().
			SetAllowMoreRetries(allow).
			SetByPassRetryCountCheck(bypass)
	}
}
