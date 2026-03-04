package builder

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	"github.com/gobeetle/fetch/internal/handler/defaults"
	mreq "github.com/gobeetle/fetch/internal/request_modifier"
	rreq "github.com/gobeetle/fetch/internal/result_request"
)

// this helps to build a http request with the use of modifiers
// it is helpful when you want to build a request with modifiers
// it is helpful for external integration without using fetch instance
func BuildRequest[Req any](modifiers ...mreq.RequestModifier[Req]) (rreq.RequestResult, *errpkg.Error) {
	return mreq.Modify(defaults.NewRequestHandler[Req](), modifiers...).Request()
}
