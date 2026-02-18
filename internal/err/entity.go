package err

import (
	"errors"
	"strings"

	rrsp "github.com/gobeetle/fetch/internal/result_response"
	"github.com/gobeetle/fetch/internal/retry"
)

type GetRetryableFunc func(rrsp.ResponseResult, Error) *retry.Retryable

type Error struct {
	status_code int
	message     []string
	response    string
	errs        []error
}

func NewError(errs ...error) *Error {
	if len(errs) == 0 {
		return nil
	}
	if len(errs) == 1 {
		// if there is only one error, ensure it is a fetch error
		return Ensure(errs[0])
	}
	return new(errs...)
}

func Ensure(err error) *Error {
	if err == nil {
		return nil
	}
	if e, ok := err.(*Error); ok {
		return e
	}
	// try to unwrap the error
	var e *Error
	if ok := errors.As(err, &e); ok {
		return &Error{
			errs:        []error{err},
			message:     e.message,
			response:    e.response,
			status_code: e.status_code,
		}
	}
	return new(err)
}

func new(err ...error) *Error {
	return &Error{
		errs: err,
	}
}

func NewErrorFromString(err_msgs ...string) *Error {
	if len(err_msgs) == 0 {
		return nil
	} else if len(err_msgs) == 1 && len(err_msgs[0]) == 0 {
		return nil
	}
	var errs []error
	for _, err := range err_msgs {
		errs = append(errs, errors.New(err))
	}
	return &Error{
		errs: errs,
	}
}

func (e *Error) WithCode(code int) *Error {
	e.status_code = code
	return e
}

func (e *Error) WithMessage(message ...string) *Error {
	e.message = message
	return e
}

func (e *Error) WithResponse(response string) *Error {
	e.response = response
	return e
}

func (e *Error) WithError(errs ...error) *Error {
	e.errs = errs
	return e
}

func (e *Error) Error() string {
	var parts []string
	if len(e.errs) > 0 {
		parts = append(parts, errors.Join(e.errs...).Error())
	}
	if len(e.message) > 0 {
		parts = append(parts, "("+strings.Join(e.message, ",")+")")
	}
	if e.response != "" {
		parts = append(parts, "["+e.response+"]")
	}
	return strings.Join(parts, " | ")
}

func (e *Error) Unwrap() []error {
	return e.errs
}

func (e *Error) StatusCode() int {
	return e.status_code
}

func (e *Error) Message() []string {
	return e.message
}
