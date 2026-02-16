package response_code

import "github.com/gobeetle/fetch/internal/enum"

func New() *ResponseCode {
	return &ResponseCode{
		StatusCodeValidation: []enum.StatusCodeFilterType{},
	}
}
