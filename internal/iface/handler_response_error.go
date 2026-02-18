package iface

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
)

type ResponseErrorHandler interface {
	ResponseComponentHandler
	WithUseErrorUnwrapper()
	WithRetryableFunc(getRetryableFuncs ...errpkg.GetRetryableFunc)
	WithError(err error)
	Error() error
}
