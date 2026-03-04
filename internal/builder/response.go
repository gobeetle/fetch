package builder

import (
	"net/http"

	errpkg "github.com/gobeetle/fetch/internal/err"
	"github.com/gobeetle/fetch/internal/handler/defaults"
	mrsp "github.com/gobeetle/fetch/internal/response_modifier"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

// this is helpful when you want to build a response with the use of modifiers
// it is helpful when you have an existing http response and
// want to build a response result with modifiers
// it is helpful for external integration without using fetch instance
func BuildResponse[Rsp any](response *http.Response, modifiers ...mrsp.ResponseModifier[Rsp]) (rrsp.ResponseResult, *errpkg.Error) {
	return mrsp.Modify(defaults.NewResponseHandler[Rsp](), modifiers...).BuildResponse(response)
}
