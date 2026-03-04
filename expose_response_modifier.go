package fetch

import (
	mrsp "github.com/gobeetle/fetch/internal/response_modifier"
)

type (
	ResponseModifier[Rsp any] = mrsp.ResponseModifier[Rsp]
)

const ()

var ()

func ModifyResponse[Rsp any](handler *ResponseHandler[Rsp], modifiers ...ResponseModifier[Rsp]) *ResponseHandler[Rsp] {
	return mrsp.Modify(handler, modifiers...)
}

func WithValidStatusCodes[Rsp any](codes ...int) ResponseModifier[Rsp] {
	return mrsp.WithValidStatusCodes[Rsp](codes...)
}

func WithRspValidStatusCodeTypes[Rsp any](code_types ...ValidCodeType) ResponseModifier[Rsp] {
	return mrsp.WithValidStatusCodeTypes[Rsp](code_types...)
}

func WithRspClearedValidStatusCodes[Rsp any]() ResponseModifier[Rsp] {
	return mrsp.WithClearedValidStatusCodes[Rsp]()
}

func WithRsp2XXAsValidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return mrsp.With2XXAsValidStatusCode[Rsp]()
}

func WithRsp3XXAsValidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return mrsp.With3XXAsValidStatusCode[Rsp]()
}

func WithRsp4XXAsValidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return mrsp.With4XXAsValidStatusCode[Rsp]()
}

func WithRsp5XXAsValidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return mrsp.With5XXAsValidStatusCode[Rsp]()
}

func WithRspAllAsValidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return mrsp.WithAllAsValidStatusCode[Rsp]()
}

func WithRspAllAsInvalidStatusCode[Rsp any]() ResponseModifier[Rsp] {
	return mrsp.WithAllAsInvalidStatusCode[Rsp]()
}

func WithRspJsonObj[Rsp any](resp_obj *Rsp) ResponseModifier[Rsp] {
	return mrsp.WithJsonObj[Rsp](resp_obj)
}

func WithRspUseErrorUnwrapper[Rsp any]() ResponseModifier[Rsp] {
	return mrsp.WithUseErrorUnwrapper[Rsp]()
}

func WithRspError[Rsp any](err error) ResponseModifier[Rsp] {
	return mrsp.WithError[Rsp](err)
}

func WithRspRetryableFunc[Rsp any](f ...GetRetryableFunc) ResponseModifier[Rsp] {
	return mrsp.WithRetryableFunc[Rsp](f...)
}
