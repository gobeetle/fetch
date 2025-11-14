package fetch

import (
	mreq "github.com/gobeetle/fetch/internal/request_modifier"
)

type (
	RequestModifier = mreq.RequestModifier
)

const ()

var (
	ModifyRequest = mreq.Modify

	WithReqJsonBody                = mreq.WithJsonBody
	WithReqRawBody                 = mreq.WithRawBody
	WithReqForm                    = mreq.WithForm
	WithReqClient                  = mreq.WithClient
	WithReqHeaderContentTypeAsJson = mreq.WithHeaderContentTypeAsJson
	WithReqHeaderContentType       = mreq.WithHeaderContentType
	WithReqHeaderAuth              = mreq.WithHeaderAuth
	WithReqHeader                  = mreq.WithHeader
	WithReqHeaderPair              = mreq.WithHeaderPair
	WithReqHeaderCleared           = mreq.WithHeaderCleared
	WithReqMethod                  = mreq.WithMethod
	WithReqMethodString            = mreq.WithMethodString
	WithReqUrl                     = mreq.WithUrl
	WithReqQueryParams             = mreq.WithQueryParams
	WithReqQueryParamPair          = mreq.WithQueryParamPair
	WithReqQueryParamsCleared      = mreq.WithQueryParamsCleared
)
