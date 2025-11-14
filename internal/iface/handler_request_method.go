package iface

import "github.com/gobeetle/fetch/internal/enum"

type RequestMethodHandler interface {
	RequestComponentHandler
	WithMethod(method enum.MethodType)
	WithMethodString(method string)
}
