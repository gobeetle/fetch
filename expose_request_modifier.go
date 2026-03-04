package fetch

import (
	"net/http"

	mreq "github.com/gobeetle/fetch/internal/request_modifier"
)

type (
	RequestModifier[Req any] = mreq.RequestModifier[Req]
)

const ()

var ()

func ModifyRequest[Req any](handler *RequestHandler[Req], modifiers ...RequestModifier[Req]) *RequestHandler[Req] {
	return mreq.Modify(handler, modifiers...)
}

func WithReqJsonBody[Req any](body *Req) RequestModifier[Req] {
	return mreq.WithJsonBody[Req](body)
}

func WithReqRawBody[Req any](body []byte) RequestModifier[Req] {
	return mreq.WithRawBody[Req](body)
}

func WithReqForm[Req any](body map[string]any) RequestModifier[Req] {
	return mreq.WithForm[Req](body)
}

func WithReqClient[Req any](client *http.Client) RequestModifier[Req] {
	return mreq.WithClient[Req](client)
}

func WithReqHeaderContentTypeAsJson[Req any]() RequestModifier[Req] {
	return mreq.WithHeaderContentTypeAsJson[Req]()
}

func WithReqHeaderContentType[Req any](content_type MediaType) RequestModifier[Req] {
	return mreq.WithHeaderContentType[Req](content_type)
}

func WithReqHeaderAuth[Req any](auth string) RequestModifier[Req] {
	return mreq.WithHeaderAuth[Req](auth)
}

func WithReqHeader[Req any](header map[string]string) RequestModifier[Req] {
	return mreq.WithHeader[Req](header)
}

func WithReqHeaderPair[Req any](key string, value string) RequestModifier[Req] {
	return mreq.WithHeaderPair[Req](key, value)
}

func WithReqHeaderCleared[Req any]() RequestModifier[Req] {
	return mreq.WithHeaderCleared[Req]()
}

func WithReqMethod[Req any](method MethodType) RequestModifier[Req] {
	return mreq.WithMethod[Req](method)
}

func WithReqMethodString[Req any](method string) RequestModifier[Req] {
	return mreq.WithMethodString[Req](method)
}

func WithReqUrl[Req any](endpoint string) RequestModifier[Req] {
	return mreq.WithUrl[Req](endpoint)
}

func WithReqQueryParams[Req any](params map[string]string) RequestModifier[Req] {
	return mreq.WithQueryParams[Req](params)
}

func WithReqQueryParamPair[Req any](key string, value string) RequestModifier[Req] {
	return mreq.WithQueryParamPair[Req](key, value)
}

func WithReqQueryParamsCleared[Req any]() RequestModifier[Req] {
	return mreq.WithQueryParamsCleared[Req]()
}
