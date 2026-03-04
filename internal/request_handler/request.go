package request_handler

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	"github.com/gobeetle/fetch/internal/iface"
	rreq "github.com/gobeetle/fetch/internal/result_request"
)

func (r *RequestHandler[Req]) Request() (rreq.RequestResult, *errpkg.Error) {
	result := rreq.New()
	for _, candidate := range []iface.RequestComponentHandler{
		r.RequestLocationHandler,
		r.RequestMethodHandler,
		r.RequestRequestHandler,
		r.RequestClientHandler,
		r.RequestHeaderHandler,
		r.RequestBodyHandler,
	} {
		if candidate != nil {
			if err := candidate.PrepareRequest(&result); err != nil {
				return result, err
			}
		}
	}
	return result, nil
}
