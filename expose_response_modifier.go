package fetch

import (
	mrsp "github.com/gobeetle/fetch/internal/response_modifier"
)

type (
	ResponseModifier = mrsp.ResponseModifier
)

const ()

var (
	ModifyResponse = mrsp.Modify

	WithRspValidStatusCodes        = mrsp.WithValidStatusCodes
	WithRspValidStatusCodeTypes    = mrsp.WithValidStatusCodeTypes
	WithRspClearedValidStatusCodes = mrsp.WithClearedValidStatusCodes
	WithRsp2XXAsValidStatusCode    = mrsp.With2XXAsValidStatusCode
	WithRsp3XXAsValidStatusCode    = mrsp.With3XXAsValidStatusCode
	WithRsp4XXAsValidStatusCode    = mrsp.With4XXAsValidStatusCode
	WithRsp5XXAsValidStatusCode    = mrsp.With5XXAsValidStatusCode
	WithRspAllAsValidStatusCode    = mrsp.WithAllAsValidStatusCode
	WithRspAllAsInvalidStatusCode  = mrsp.WithAllAsInvalidStatusCode
	WithRspJsonObj                 = mrsp.WithJsonObj
	WithRspJsonObjectWrapper       = mrsp.WithJsonObjectWrapper
	WithRspUseErrorUnwrapper       = mrsp.WithUseErrorUnwrapper
	WithRspError                   = mrsp.WithError
)
