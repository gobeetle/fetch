package request_modifier

import (
	"net/http"

	"github.com/gobeetle/fetch/internal/enum"
	hreq "github.com/gobeetle/fetch/internal/request_handler"
)

/*
this is a modifier that set the request body to a json object
*/
func WithJsonBody[Req any](body *Req) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestBodyHandler.WithJsonBody(body)
		return r
	}
}

/*
this is a modifier that set the request body to a raw byte array
*/
func WithRawBody[Req any](body []byte) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestBodyHandler.WithRawBody(body)
		return r
	}
}

/*
this is a modifier that set the request body to a form
*/
func WithForm[Req any](body map[string]any) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestBodyHandler.WithForm(body)
		return r
	}
}

/*
this is a modifier that set the http client
*/
func WithClient[Req any](client *http.Client) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestClientHandler.WithClient(client)
		return r
	}
}

/*
this is a modifier that set the header content type to json
*/
func WithHeaderContentTypeAsJson[Req any]() RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestHeaderHandler.WithHeaderContentTypeAsJson()
		return r
	}
}

/*
this is a modifier that set the header content type to a specific content type
*/
func WithHeaderContentType[Req any](content_type enum.MediaType) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestHeaderHandler.WithHeaderContentType(content_type)
		return r
	}
}

/*
this is a modifier that set the value to the Authorization key in the header
*/
func WithHeaderAuth[Req any](token string) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestHeaderHandler.WithHeaderAuth(token)
		return r
	}
}

/*
this is a modifier that set the header
*/
func WithHeader[Req any](header map[string]string) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestHeaderHandler.WithHeader(header)
		return r
	}
}

/*
this is a modifier that set the header
*/
func WithHeaderPair[Req any](key string, value string) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestHeaderHandler.WithHeaderPair(key, value)
		return r
	}
}

/*
this is a modifier that clear the header
*/
func WithHeaderCleared[Req any]() RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestHeaderHandler.WithHeaderCleared()
		return r
	}
}

/*
this is a modifier that set the method of the request
*/
func WithMethod[Req any](method enum.MethodType) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestMethodHandler.WithMethod(method)
		return r
	}
}

/*
this is a modifier that set the method of the request
*/
func WithMethodString[Req any](method string) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestMethodHandler.WithMethodString(method)
		return r
	}
}

/*
req_opt_modifier is a function that set request url
*/
func WithUrl[Req any](endpoint string) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestLocationHandler.WithUrl(endpoint)
		return r
	}
}

/*
req_opt_modifier is a function that set query parameters
*/
func WithQueryParams[Req any](params map[string]string) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestLocationHandler.WithQueryParams(params)
		return r
	}
}

/*
this is a modifier that add a query parameter to the request
*/
func WithQueryParamPair[Req any](key string, value string) RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestLocationHandler.WithQueryParamPair(key, value)
		return r
	}
}

/*
this is a modifier that clear the query parameters of the request
*/
func WithQueryParamsCleared[Req any]() RequestModifier[Req] {
	return func(r *hreq.RequestHandler[Req]) *hreq.RequestHandler[Req] {
		r.RequestLocationHandler.WithQueryParamsCleared()
		return r
	}
}
