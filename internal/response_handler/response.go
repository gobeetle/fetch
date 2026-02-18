package response_handler

import (
	"fmt"
	"net/http"

	errpkg "github.com/gobeetle/fetch/internal/err"
	"github.com/gobeetle/fetch/internal/iface"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

func (r *ResponseHandler) BuildResponse(response *http.Response) (rrsp.ResponseResult, *errpkg.Error) {
	result := rrsp.New()
	result.Response = response
	if result.Response == nil {
		return rrsp.ResponseResult{}, errpkg.NewError(
			fmt.Errorf("response cannot be nil"),
		)
	}
	if err := r.Response(&result); err != nil {
		return result, err
	}
	return result, nil
}

func (r *ResponseHandler) Response(result *rrsp.ResponseResult) (err *errpkg.Error) {
	defer func() {
		if r.ResponseErrorHandler != nil {
			f := r.ResponseErrorHandler.GetRetryableFunc()
			if f != nil {
				// if we have a retryable function, we should use it
				errForRetryable := err
				if errForRetryable == nil {
					errForRetryable = &errpkg.Error{}
				}
				retryable := f(*result, *errForRetryable)
				if retryable != nil {
					result.SetAllowMoreRetries(retryable.AllowMoreRetries())
					result.SetByPassRetryCountCheck(retryable.ByPassRetryCountCheck())
				}
			}
			if result.HttpError != nil {
				// if we have http error, we should override the error with the http error
				err = errpkg.NewError(result.HttpError)
			}
		}
	}()

	// result.Response = response
	if result.Response == nil {
		return errpkg.NewError(
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
			candidate.Prepare(result)
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
			if err := candidate.ValidateResponse(result); err != nil {
				return errpkg.
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
			if err := candidate.PrepareResponse(result); err != nil {
				return errpkg.
					NewError(err).
					WithCode(result.StatusCode).
					WithResponse(string(result.RespBytes))
			}
		}
	}

	return nil
}
