package response_body

import errpkg "github.com/gobeetle/fetch/internal/err"

func (r *ResponseBody[Rsp]) GetRetryableFunc() errpkg.GetRetryableFunc {
	return nil
}
