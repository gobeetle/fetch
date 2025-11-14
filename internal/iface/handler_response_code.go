package iface

import "github.com/gobeetle/fetch/internal/enum"

type ResponseCodeHandler interface {
	ResponseComponentHandler
	WithValidStatusCodes(codes ...int)
	WithClearedValidStatusCodes()
	WithValidStatusCodeTypes(codeTypes ...enum.StatusCodeFilterType)
	With2XXAsValidStatusCode()
	With3XXAsValidStatusCode()
	With4XXAsValidStatusCode()
	With5XXAsValidStatusCode()
	WithAllAsValidStatusCode()
	WithAllAsInvalidStatusCode()
}
