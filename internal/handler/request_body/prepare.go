package request_body

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"

	"github.com/gobeetle/fetch/internal/enum"
	errpkg "github.com/gobeetle/fetch/internal/err"
	"github.com/gobeetle/fetch/internal/handler/request_header"
	rreq "github.com/gobeetle/fetch/internal/result_request"
)

func (r *RequestBody) PrepareRequest(result *rreq.RequestResult) *errpkg.Error {
	req := result.Request
	if r.RawBody != nil {
		req.Body = io.NopCloser(bytes.NewReader(r.RawBody))
	}
	if r.FormBody != nil {
		request_header.SetHeaderContentType(req, enum.EnumMediaType.MediaTypeForm)
		form := url.Values{}
		for key, value := range r.FormBody {
			form.Add(key, fmt.Sprintf("%v", value))
		}

		req.Body = io.NopCloser(strings.NewReader(form.Encode()))
	}
	if r.JsonBody != nil {
		request_header.SetHeaderContentType(req, enum.EnumMediaType.MediaTypeJson)
		var raw_body []byte
		switch v := r.JsonBody.(type) {
		case []byte:
			raw_body = v
		case string:
			raw_body = []byte(v)
		default:
			r, err := json.Marshal(r.JsonBody)
			if err != nil {
				return errpkg.NewError(
					fmt.Errorf("prepare request body: %v", err),
				)
			}
			raw_body = r
		}
		req.Body = io.NopCloser(bytes.NewReader(raw_body))
	}
	return nil
}
