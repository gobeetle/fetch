package response_code

import (
	"fmt"
	"slices"
	"strings"

	"github.com/gobeetle/fetch/internal/enum"
	errpkg "github.com/gobeetle/fetch/internal/err"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

func (r *ResponseCode) ValidateResponse(result *rrsp.ResponseResult) *errpkg.Error {
	statusCode := result.StatusCode
	failed := errpkg.NewError(
		fmt.Errorf(
			"status code %d is not considered valid, filter: %s",
			statusCode, strings.Join(r.StatusCodeFilter(), ", "),
		),
	)

	if len(r.StatusCodeValidation) == 0 && len(r.ValidStatusCodes) == 0 {
		// if no validation set, skip validation
		return nil
	}

	if slices.Contains(r.StatusCodeValidation, enum.EnumStatusCodeFilterType.AllInvalid) {
		// if all invalid, return error
		return failed
	} else if slices.Contains(r.StatusCodeValidation, enum.EnumStatusCodeFilterType.AllValid) {
		// if all valid, skip validation
		return nil
	}

	// check if any of the valid status code types match
	for _, codeType := range r.StatusCodeValidation {
		switch codeType {
		case enum.EnumStatusCodeFilterType.AllValid:
			return nil
		case enum.EnumStatusCodeFilterType.Include2XX:
			if statusCode >= 200 && statusCode < 300 {
				return nil
			}
		case enum.EnumStatusCodeFilterType.Include3XX:
			if statusCode >= 300 && statusCode < 400 {
				return nil
			}
		case enum.EnumStatusCodeFilterType.Include4XX:
			if statusCode >= 400 && statusCode < 500 {
				return nil
			}
		case enum.EnumStatusCodeFilterType.Include5XX:
			if statusCode >= 500 && statusCode < 600 {
				return nil
			}
		}
	}

	// check specific status codes if provided
	if len(r.ValidStatusCodes) > 0 {
		if slices.Contains(r.ValidStatusCodes, statusCode) {
			return nil
		}
	}

	// skip validation if status code is 0 and it still fails at this point
	if statusCode == 0 {
		return nil
	}

	return failed
}
func (r *ResponseCode) Validate() error {
	return nil
}
