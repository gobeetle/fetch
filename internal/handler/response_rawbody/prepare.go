package response_rawbody

import (
	"io"

	errpkg "github.com/gobeetle/fetch/internal/err"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

func (r *ResponseRawBody) Prepare(result *rrsp.ResponseResult) {
	//if any of body string or body bytes is set, use it instead of reading from response
	if len(r.BodyBytes) > 0 {
		result.SetRespBytes(r.BodyBytes)
		return
	}
	if resp := result.GetResponse(); resp != nil {
		defer resp.Body.Close()
		rawResponse, _ := io.ReadAll(resp.Body)
		result.SetRespBytes(rawResponse)
	}
}

func (r *ResponseRawBody) PrepareResponse(result *rrsp.ResponseResult) *errpkg.Error { return nil }
