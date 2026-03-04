package fetch

import "github.com/gobeetle/fetch/internal/handler/request_location"

type (
	RequestLocation = request_location.RequestLocation
)

const ()

var ()

func NewRequestLocation() *RequestLocation {
	return request_location.New()
}
