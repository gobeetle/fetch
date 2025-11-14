package response_handler

import (
	"github.com/gobeetle/fetch/internal/iface"
)

type ResponseHandler struct {
	/*
	   error
	*/
	iface.ResponseErrorHandler

	/*
	   code
	*/
	iface.ResponseCodeHandler

	/*
	   raw body
	*/
	iface.ResponseRawBodyHandler

	/*
	   body
	*/
	iface.ResponseBodyHandler
}

func (r *ResponseHandler) Error() error {
	if r.ResponseErrorHandler == nil {
		return nil
	}
	ierr, ok := r.ResponseErrorHandler.(interface{ Error() error })
	if !ok {
		return nil
	}
	return ierr.Error()
}
