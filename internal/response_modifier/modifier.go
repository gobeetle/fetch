package response_modifier

import (
	"github.com/gobeetle/fetch/internal/enum"
	errpkg "github.com/gobeetle/fetch/internal/err"
	hrsp "github.com/gobeetle/fetch/internal/response_handler"
)

/*
rsp_opt_modifier is a function that set the validation of the response status codes
*/
func WithValidStatusCodes[Rsp any](codes ...int) ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseCodeHandler.WithValidStatusCodes(codes...)
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func WithValidStatusCodeTypes[Rsp any](code_types ...enum.StatusCodeFilterType) ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseCodeHandler.WithValidStatusCodeTypes(code_types...)
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func WithClearedValidStatusCodes[Rsp any]() ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseCodeHandler.WithClearedValidStatusCodes()
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func With2XXAsValidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseCodeHandler.With2XXAsValidStatusCode()
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func With3XXAsValidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseCodeHandler.With3XXAsValidStatusCode()
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func With4XXAsValidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseCodeHandler.With4XXAsValidStatusCode()
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func With5XXAsValidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseCodeHandler.With5XXAsValidStatusCode()
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func WithAllAsValidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseCodeHandler.WithAllAsValidStatusCode()
		return r
	}
}

func WithAllAsInvalidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseCodeHandler.WithAllAsInvalidStatusCode()
		return r
	}
}

/*
this is a modifier that set the response object
*/
func WithJsonObj[Rsp any](resp_obj *Rsp) ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseBodyHandler.WithJsonObj(resp_obj)
		return r
	}
}

/*
this is a modifier that set the response object with a wrapper
*/
func WithUseErrorUnwrapper[Rsp any]() ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseErrorHandler.WithUseErrorUnwrapper()
		return r
	}
}

/*
this is a modifier that set the error
*/
func WithError[Rsp any](err error) ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseErrorHandler.WithError(err)
		return r
	}
}

func WithRetryableFunc[Rsp any](f ...errpkg.GetRetryableFunc) ResponseModifier[Rsp] {
	return func(r *hrsp.ResponseHandler[Rsp]) *hrsp.ResponseHandler[Rsp] {
		r.ResponseErrorHandler.WithRetryableFunc(f...)
		return r
	}
}
