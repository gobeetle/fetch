package iface

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
)

type ResponseRetryFuncHandler interface {
	ResponseComponentHandler
	WithGetRetryableFunc(getRetryableFunc errpkg.GetRetryableFunc)
}
