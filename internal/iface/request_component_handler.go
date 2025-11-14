package iface

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
	rreq "github.com/gobeetle/fetch/internal/result_request"
)

type RequestComponentHandler interface {
	Validator
	PrepareRequest(*rreq.RequestResult) *errpkg.Error
}
