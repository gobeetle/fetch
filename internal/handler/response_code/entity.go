package response_code

import (
	"github.com/gobeetle/fetch/internal/enum"
	"github.com/gobeetle/fetch/internal/iface"
)

type ResponseCode struct {
	StatusCode           *int
	ValidStatusCodes     []int                       // status codes to validate, if not provided, all status codes are valid
	StatusCodeValidation []enum.StatusCodeFilterType // additional status codes to validate
}

var (
	_ iface.ResponseCodeHandler = (*ResponseCode)(nil)
)
