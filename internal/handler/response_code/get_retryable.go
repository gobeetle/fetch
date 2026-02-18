package response_code

import errpkg "github.com/gobeetle/fetch/internal/err"

func (r *ResponseCode) GetRetryableFunc() errpkg.GetRetryableFunc {
	return nil
}
