package iface

type ResponseErrorHandler interface {
	ResponseComponentHandler
	WithUseErrorUnwrapper()
	WithError(err error)
	Error() error
}
