package retry

import "time"

func New[A, B any]() *Retry[A, B] {
	return &Retry[A, B]{
		count: DefaultRetryCount,
		delay: DefaultRetryDelay,
	}
}

// WithRetry sets the retry count and delay
// if the retry count is set to InfiniteRetries (-1), it will keep retrying until success
// if count is 0, it will not retry after the first attempt
// if count is positive, it will retry up to count times
// delay is the duration to wait before retrying after each failed attempt
func (r *Retry[A, B]) WithRetry(count int, delay time.Duration) *Retry[A, B] {
	r.count = count
	r.delay = delay
	return r
}
