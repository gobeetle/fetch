package fetch

import (
	"github.com/gobeetle/fetch/internal/fetch"
)

type (
	Fetch[Req, Rsp any] = fetch.Fetch[Req, Rsp]
)

const ()

var ()

func New[Req, Rsp any]() *fetch.Fetch[Req, Rsp] {
	return fetch.New[Req, Rsp]()
}
