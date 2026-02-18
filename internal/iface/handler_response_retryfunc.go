package iface

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
)

type ResponseRetryFuncHandler interface {
	ResponseComponentHandler
	WithGetRetryableFunc(getRetryableFuncs ...errpkg.GetRetryableFunc)
}
