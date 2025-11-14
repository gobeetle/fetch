package fetch

import "github.com/gobeetle/fetch/internal/err"

type (
	Error = err.Error
)

const ()

var (
	NewError           = err.NewError
	NewErrorFromString = err.NewErrorFromString
)
