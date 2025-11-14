package retry

import (
	"time"
)

type Retry[A, B any] struct {
	count int
	delay time.Duration
}

const (
	// retry count, 0 means no retry, it will not retry after the first attempt
	// -1 means infinite retries
	DefaultRetryCount = 0

	// retry delay, it will wait for this duration before retrying after each failed attempt
	DefaultRetryDelay = 0

	// InfiniteRetries can be used with WithRetry to specify infinite retries
	InfiniteRetries = -1
)

type DoFunc[A, B any] func(input A) (B, error)

// do executes the provided function with retry logic.
// if the retry count is set to InfiniteRetries (-1), it will keep retrying until success.
func (r *Retry[A, B]) Do(fn DoFunc[A, B], input A) (B, error) {
	var (
		err    error
		result B
		i      = 0
	)

	for {
		// if the retry count is not infinite and we've exceeded the retry count, break
		if r.count != InfiniteRetries && i > r.count {
			break
		}

		result, err = fn(input)
		if err == nil {
			return result, nil
		}

		// if this was the last attempt, break and return the error
		if r.count != InfiniteRetries && i == r.count {
			break
		}

		if r.delay > 0 {
			// if delay is set, wait for the delay duration before retrying
			time.Sleep(r.delay)
		}
		i++
	}
	return result, err
}
