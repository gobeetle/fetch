package response_code

import (
	"fmt"

	"github.com/gobeetle/fetch/internal/enum"
)

func (r *ResponseCode) WithStatusCode(status_code int) {
	r.StatusCode = &status_code
}

func (r *ResponseCode) WithValidStatusCodeTypes(code_types ...enum.StatusCodeFilterType) {
	r.StatusCodeValidation = code_types
}

func (r *ResponseCode) WithValidStatusCodes(codes ...int) {
	r.ValidStatusCodes = codes
}

func (r *ResponseCode) WithClearedValidStatusCodes() {
	r.ValidStatusCodes = []int{}
}

func (r *ResponseCode) With2XXAsValidStatusCode() {
	r.WithValidStatusCodeTypes(enum.EnumStatusCodeFilterType.Include2XX)
}

func (r *ResponseCode) With3XXAsValidStatusCode() {
	r.WithValidStatusCodeTypes(enum.EnumStatusCodeFilterType.Include3XX)
}

func (r *ResponseCode) With4XXAsValidStatusCode() {
	r.WithValidStatusCodeTypes(enum.EnumStatusCodeFilterType.Include4XX)
}

func (r *ResponseCode) With5XXAsValidStatusCode() {
	r.WithValidStatusCodeTypes(enum.EnumStatusCodeFilterType.Include5XX)
}

func (r *ResponseCode) WithAllAsValidStatusCode() {
	r.WithClearedValidStatusCodes()
	r.WithValidStatusCodeTypes(enum.EnumStatusCodeFilterType.AllValid)
}

func (r *ResponseCode) WithAllAsInvalidStatusCode() {
	r.WithClearedValidStatusCodes()
	r.WithValidStatusCodeTypes(enum.EnumStatusCodeFilterType.AllInvalid)
}

func (r *ResponseCode) StatusCodeFilter() []string {
	result := []string{}
	for _, code := range r.ValidStatusCodes {
		result = append(result, fmt.Sprintf("%d", code))
	}
	for _, codeType := range r.StatusCodeValidation {
		switch codeType {
		case enum.EnumStatusCodeFilterType.Include2XX:
			result = append(result, "2XX")
		case enum.EnumStatusCodeFilterType.Include3XX:
			result = append(result, "3XX")
		case enum.EnumStatusCodeFilterType.Include4XX:
			result = append(result, "4XX")
		case enum.EnumStatusCodeFilterType.Include5XX:
			result = append(result, "5XX")
		case enum.EnumStatusCodeFilterType.AllValid:
			result = append(result, "AllValid")
		case enum.EnumStatusCodeFilterType.AllInvalid:
			result = append(result, "AllInvalid")
		}
	}
	if len(result) == 0 {
		result = append(result, "N/A")
	}
	return result
}
