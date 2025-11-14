package response_modifier

import (
	"github.com/gobeetle/fetch/internal/enum"
	hrsp "github.com/gobeetle/fetch/internal/response_handler"
)

/*
rsp_opt_modifier is a function that set the validation of the response status codes
*/
func WithValidStatusCodes(codes ...int) ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseCodeHandler.WithValidStatusCodes(codes...)
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func WithValidStatusCodeTypes(code_types ...enum.StatusCodeFilterType) ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseCodeHandler.WithValidStatusCodeTypes(code_types...)
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func WithClearedValidStatusCodes() ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseCodeHandler.WithClearedValidStatusCodes()
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func With2XXAsValidStatusCode() ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseCodeHandler.With2XXAsValidStatusCode()
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func With3XXAsValidStatusCode() ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseCodeHandler.With3XXAsValidStatusCode()
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func With4XXAsValidStatusCode() ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseCodeHandler.With4XXAsValidStatusCode()
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func With5XXAsValidStatusCode() ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseCodeHandler.With5XXAsValidStatusCode()
		return r
	}
}

/*
this is a modifier that set the validation of the response status code types
*/
func WithAllAsValidStatusCode() ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseCodeHandler.WithAllAsValidStatusCode()
		return r
	}
}

func WithAllAsInvalidStatusCode() ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseCodeHandler.WithAllAsInvalidStatusCode()
		return r
	}
}

/*
this is a modifier that set the response object
*/
func WithJsonObj(resp_obj any) ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseBodyHandler.WithJsonObj(resp_obj)
		return r
	}
}

/*
this is a modifier that set the response object with a wrapper
*/
func WithJsonObjectWrapper(wrapper any, data_field_name string, data_obj any) ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseBodyHandler.WithJsonObjectWrapper(wrapper, data_field_name, data_obj)
		return r
	}
}

/*
this is a modifier that set the response object with a wrapper
*/
func WithUseErrorUnwrapper() ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseErrorHandler.WithUseErrorUnwrapper()
		return r
	}
}

/*
this is a modifier that set the error
*/
func WithError(err error) ResponseModifier {
	return func(r *hrsp.ResponseHandler) *hrsp.ResponseHandler {
		r.ResponseErrorHandler.WithError(err)
		return r
	}
}
