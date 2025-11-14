package response_handler

import (
	"fmt"
	"net/http"

	errpkg "github.com/gobeetle/fetch/internal/err"
	"github.com/gobeetle/fetch/internal/iface"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

func (r *ResponseHandler) Response(response *http.Response) (rrsp.ResponseResult, *errpkg.Error) {
	result := rrsp.New()
	result.Response = response
	if result.Response == nil {
		return rrsp.ResponseResult{}, errpkg.NewError(
			fmt.Errorf("response cannot be nil"),
		)
	}

	// fast prepare
	for _, candidate := range []iface.ResponseComponentHandler{
		r.ResponseRawBodyHandler, // first get the raw response into the result
		r.ResponseCodeHandler,    // then get the code into the result
		r.ResponseBodyHandler,    // n/a atm
		r.ResponseErrorHandler,   // n/a atm
	} {
		if candidate != nil {
			candidate.Prepare(&result)
		}
	}

	// fast validate
	for _, candidate := range []iface.ResponseComponentHandler{
		r.ResponseRawBodyHandler, // n/a atm
		r.ResponseCodeHandler,    // validate the code
		r.ResponseBodyHandler,    // n/a atm
		r.ResponseErrorHandler,   // n/a atm
	} {
		if candidate != nil {
			if err := candidate.ValidateResponse(&result); err != nil {
				return result, errpkg.
					NewError(err).
					WithCode(result.StatusCode).
					WithResponse(string(result.RespBytes))
			}
		}
	}

	// full prepare
	for _, candidate := range []iface.ResponseComponentHandler{
		r.ResponseRawBodyHandler, // n/a atm
		r.ResponseCodeHandler,    // n/a atm
		r.ResponseBodyHandler,    // prepare the body
		r.ResponseErrorHandler,   // prepare the error
	} {
		if candidate != nil {
			if err := candidate.PrepareResponse(&result); err != nil {
				return result, errpkg.
					NewError(err).
					WithCode(result.StatusCode).
					WithResponse(string(result.RespBytes))
			}
		}
	}

	return result, nil
}
