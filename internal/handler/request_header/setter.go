package request_header

import (
	"net/http"

	"github.com/gobeetle/fetch/internal/enum"
)

func (r *RequestHeader) WithHeaderContentTypeAsJson() {
	r.HeaderContentType = enum.EnumMediaType.MediaTypeJson
}

func (r *RequestHeader) WithHeaderContentType(content_type enum.MediaType) {
	r.HeaderContentType = content_type
}

func (r *RequestHeader) WithHeaderAuth(auth string) {
	r.HeaderAuth = auth
}

func (r *RequestHeader) WithHeader(header map[string]string) {
	r.Header = header
}

func (r *RequestHeader) WithHeaderPair(key string, value string) {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = value
}

func (r *RequestHeader) WithHeaderCleared() {
	r.Header = make(map[string]string)
}

func SetHeaderContentType(req *http.Request, content_type enum.MediaType) {
	req.Header.Set("Content-Type", string(content_type))
}
