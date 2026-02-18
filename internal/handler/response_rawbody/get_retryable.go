package response_rawbody

import errpkg "github.com/gobeetle/fetch/internal/err"

func (r *ResponseRawBody) GetRetryableFunc() errpkg.GetRetryableFunc {
	return nil
}
