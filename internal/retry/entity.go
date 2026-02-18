package retry

import (
	"time"
)

type Retry[A any, B IRetryable] struct {
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

type DoFunc[A any, B IRetryable] func(input A) (B, error)

// do executes the provided function with retry logic.
// if the retry count is set to InfiniteRetries (-1), it will keep retrying until success.
func (r *Retry[A, B]) Do(fn DoFunc[A, B], input A) (B, error) {
	var (
		err                       error
		allow_more_retry          bool
		by_pass_retry_count_check bool
		result                    B
		i                         = 0
	)

	for {
		// if the retry count is not infinite and we've exceeded the retry count, break
		if !by_pass_retry_count_check && r.count != InfiniteRetries && i > r.count {
			break
		}

		result, err = fn(input)
		by_pass_retry_count_check = result.ByPassRetryCountCheck()

		if allow_more_retry = result.AllowMoreRetries(); !allow_more_retry {
			// if we don't allow retry, return the result and error immediately
			break
		}

		if err == nil {
			// if there's no error, return the result
			return result, nil
		}

		if !by_pass_retry_count_check {
			// if by pass retry count check is not set, check retry count
			// if this was the last attempt, break and return the error
			if r.count != InfiniteRetries && i == r.count {
				break
			}
		}

		if r.delay > 0 {
			// if delay is set, wait for the delay duration before retrying
			time.Sleep(r.delay)
		}
		i++
	}
	return result, err
}

// more entities

type IRetryable interface {
	AllowMoreRetries() bool
	ByPassRetryCountCheck() bool
}

type Retryable struct {
	allowMoreRetries     bool
	byPassRetryCountFlag bool
}

func NewRetryable() *Retryable {
	return &Retryable{
		allowMoreRetries:     true,
		byPassRetryCountFlag: false,
	}
}

func (r *Retryable) AllowMoreRetries() bool {
	return r.allowMoreRetries
}

func (r *Retryable) ByPassRetryCountCheck() bool {
	return r.byPassRetryCountFlag
}

func (r *Retryable) SetAllowMoreRetries(allow bool) *Retryable {
	r.allowMoreRetries = allow
	return r
}

func (r *Retryable) SetByPassRetryCountCheck(byPass bool) *Retryable {
	r.byPassRetryCountFlag = byPass
	return r
}

var _ IRetryable = (*Retryable)(nil)
