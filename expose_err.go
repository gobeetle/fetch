package fetch

import (
	errpkg "github.com/gobeetle/fetch/internal/err"
)

type (
	Error            = errpkg.Error
	GetRetryableFunc = errpkg.GetRetryableFunc
)

const ()

var (
	NewError           = errpkg.NewError
	NewErrorFromString = errpkg.NewErrorFromString
)
