package response_err

import (
	"encoding/json"
	"fmt"

	errpkg "github.com/gobeetle/fetch/internal/err"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

type errorEnvelope struct {
	Error  json.RawMessage   `json:"error,omitempty"`
	Errors []json.RawMessage `json:"errors,omitempty"`
}

func (r *ResponseErr) Prepare(_ *rrsp.ResponseResult) {}
func (r *ResponseErr) PrepareResponse(result *rrsp.ResponseResult) *errpkg.Error {
	if !r.UseErrorUnwrapper {
		return nil
	}
	rawResponse := result.RespBytes
	if len(rawResponse) == 0 {
		return nil
	}
	// basically what this should do is, try unmarshal into the wrapper
	var envelope errorEnvelope
	if err := json.Unmarshal(rawResponse, &envelope); err != nil {
		return errpkg.NewError(
			fmt.Errorf("unmarshal response: %w - '%v'", err, string(rawResponse)),
		)
	}
	var decoded_errors []string
	if len(envelope.Error) > 0 {
		decoded_errors = append(decoded_errors, string(envelope.Error))
	}
	if len(envelope.Errors) > 0 {
		for _, err := range envelope.Errors {
			decoded_errors = append(decoded_errors, string(err))
		}
	}
	if len(decoded_errors) > 0 {
		return errpkg.NewErrorFromString(
			decoded_errors...,
		)
	}
	return nil
}
