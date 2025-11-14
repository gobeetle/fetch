package request_modifier

import (
	"net/http"

	"github.com/gobeetle/fetch/internal/enum"
	hreq "github.com/gobeetle/fetch/internal/request_handler"
)

/*
this is a modifier that set the request body to a json object
*/
func WithJsonBody(body any) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestBodyHandler.WithJsonBody(body)
		return r
	}
}

/*
this is a modifier that set the request body to a raw byte array
*/
func WithRawBody(body []byte) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestBodyHandler.WithRawBody(body)
		return r
	}
}

/*
this is a modifier that set the request body to a form
*/
func WithForm(body map[string]any) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestBodyHandler.WithForm(body)
		return r
	}
}

/*
this is a modifier that set the http client
*/
func WithClient(client *http.Client) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestClientHandler.WithClient(client)
		return r
	}
}

/*
this is a modifier that set the header content type to json
*/
func WithHeaderContentTypeAsJson() RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestHeaderHandler.WithHeaderContentTypeAsJson()
		return r
	}
}

/*
this is a modifier that set the header content type to a specific content type
*/
func WithHeaderContentType(content_type enum.MediaType) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestHeaderHandler.WithHeaderContentType(content_type)
		return r
	}
}

/*
this is a modifier that set the value to the Authorization key in the header
*/
func WithHeaderAuth(token string) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestHeaderHandler.WithHeaderAuth(token)
		return r
	}
}

/*
this is a modifier that set the header
*/
func WithHeader(header map[string]string) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestHeaderHandler.WithHeader(header)
		return r
	}
}

/*
this is a modifier that set the header
*/
func WithHeaderPair(key string, value string) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestHeaderHandler.WithHeaderPair(key, value)
		return r
	}
}

/*
this is a modifier that clear the header
*/
func WithHeaderCleared() RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestHeaderHandler.WithHeaderCleared()
		return r
	}
}

/*
this is a modifier that set the method of the request
*/
func WithMethod(method enum.MethodType) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestMethodHandler.WithMethod(method)
		return r
	}
}

/*
this is a modifier that set the method of the request
*/
func WithMethodString(method string) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestMethodHandler.WithMethodString(method)
		return r
	}
}

/*
req_opt_modifier is a function that set request url
*/
func WithUrl(endpoint string) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestLocationHandler.WithUrl(endpoint)
		return r
	}
}

/*
req_opt_modifier is a function that set query parameters
*/
func WithQueryParams(params map[string]string) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestLocationHandler.WithQueryParams(params)
		return r
	}
}

/*
this is a modifier that add a query parameter to the request
*/
func WithQueryParamPair(key string, value string) RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestLocationHandler.WithQueryParamPair(key, value)
		return r
	}
}

/*
this is a modifier that clear the query parameters of the request
*/
func WithQueryParamsCleared() RequestModifier {
	return func(r *hreq.RequestHandler) *hreq.RequestHandler {
		r.RequestLocationHandler.WithQueryParamsCleared()
		return r
	}
}
