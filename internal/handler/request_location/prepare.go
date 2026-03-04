package request_location

import (
	"fmt"
	"net/url"

	errpkg "github.com/gobeetle/fetch/internal/err"
	rreq "github.com/gobeetle/fetch/internal/result_request"
)

func (r *RequestLocation) PrepareRequest(result *rreq.RequestResult) *errpkg.Error {
	// Ensure the URL is valid by parsing it
	reqURL, err := url.Parse(r.URL)
	if err != nil {
		return errpkg.NewError(
			fmt.Errorf("parse URL: %w", err),
		)
	}

	// set url to result
	result.SetURL(reqURL.String())

	// Add query parameters to the URL if they exist
	if len(r.QueryParams) > 0 {
		query := reqURL.Query()
		for key, value := range r.QueryParams {
			query.Set(key, value)
		}
		reqURL.RawQuery = query.Encode()
		r.URL = reqURL.String()
	}

	return nil
}
