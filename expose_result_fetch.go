package fetch

import (
	rfetch "github.com/gobeetle/fetch/internal/result_fetch"
)

type (
	FetchResult = rfetch.Result
)

const ()

var ()

func NewFetchResult() FetchResult {
	return rfetch.New()
}
