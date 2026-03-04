package fetch

import (
	"net/http"

	"github.com/gobeetle/fetch/internal/builder"
)

type ()

const ()

var ()

func BuildRequest[Req any](modifiers ...RequestModifier[Req]) (RequestResult, *Error) {
	return builder.BuildRequest(modifiers...)
}

func BuildResponse[Rsp any](response *http.Response, modifiers ...ResponseModifier[Rsp]) (ResponseResult, *Error) {
	return builder.BuildResponse(response, modifiers...)
}
