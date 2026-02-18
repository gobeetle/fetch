package response_body

import errpkg "github.com/gobeetle/fetch/internal/err"

func (r *ResponseBody) GetRetryableFunc() errpkg.GetRetryableFunc {
	return nil
}
