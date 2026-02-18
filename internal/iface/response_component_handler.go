package iface

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

type ResponseComponentHandler interface {
	Validator
	Prepare(*rrsp.ResponseResult)                       // prepare process that does not return error
	PrepareResponse(*rrsp.ResponseResult) *errpkg.Error // prepare that returns error
	ValidateResponse(*rrsp.ResponseResult) *errpkg.Error
	GetRetryableFunc() errpkg.GetRetryableFunc
}
