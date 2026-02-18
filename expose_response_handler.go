package fetch

import (
	"net/http"

	errpkg "github.com/gobeetle/fetch/internal/err"
	hrsp "github.com/gobeetle/fetch/internal/response_handler"
)

type (
	ResponseHandler = hrsp.ResponseHandler
)

const ()

var (
	// this is helpful when you want to build a response with the use of modifiers
	// it is helpful when you have an existing http response and
	// want to build a response result with modifiers
	// it is helpful for external integration without using fetch instance
	BuildResponse = func(response *http.Response, modifiers ...ResponseModifier) (ResponseResult, *errpkg.Error) {
		return ModifyResponse(NewResponseHandler(), modifiers...).BuildResponse(response)
	}
)
