package request_method

import (
	"github.com/gobeetle/fetch/internal/enum"
	"github.com/gobeetle/fetch/internal/iface"
)

type RequestMethod struct {
	Method enum.MethodType
}

var (
	_ iface.RequestMethodHandler = (*RequestMethod)(nil)
)
