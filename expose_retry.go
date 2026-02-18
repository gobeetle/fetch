package fetch

import (
	"github.com/gobeetle/fetch/internal/retry"
)

type (
	Retryable = retry.Retryable
)

const ()

var (
	// NewRetryable creates a new Retryable instance
	NewRetryable = retry.NewRetryable
)
