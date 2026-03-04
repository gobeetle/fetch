package response_body

import (
	"encoding/json"
	"fmt"

	errpkg "github.com/gobeetle/fetch/internal/err"
	rrsp "github.com/gobeetle/fetch/internal/result_response"
)

func (r *ResponseBody[Rsp]) Prepare(_ *rrsp.ResponseResult) {}

// func (r *ResponseBody[Rsp]) PrepareResponse(result *rrsp.ResponseResult) *errpkg.Error {
// 	rawResponse := result.RespBytes
// 	if r.Body != nil {
// 		// Unmarshal into the wrapper first
// 		wrapperValue := reflect.ValueOf(r.Body)
// 		if wrapperValue.Kind() == reflect.Ptr {
// 			wrapperValue = wrapperValue.Elem()
// 		}
// 		if err := json.Unmarshal(rawResponse, r.Body); err != nil {
// 			return errpkg.NewError(
// 				fmt.Errorf("unmarshal response: %w - '%v'", err, string(rawResponse)),
// 			)
// 		}
// 		// Now unmarshal the data field into the correct type
// 		wrapperField := wrapperValue.FieldByName(r.WrapperDataFieldName)
// 		if wrapperField.IsValid() && wrapperField.CanInterface() {
// 			dataBytes, err := json.Marshal(wrapperField.Interface())
// 			if err != nil {
// 				return errpkg.NewError(
// 					fmt.Errorf("marshal nested data: %w - '%v'", err, string(rawResponse)),
// 				)
// 			}
// 			if err := json.Unmarshal(dataBytes, wrapperField.Addr().Interface()); err != nil {
// 				return errpkg.NewError(
// 					fmt.Errorf("unmarshal nested data: %w - '%v'", err, string(dataBytes)),
// 				)
// 			}
// 		}
// 	}
// 	return nil
// }

func (r *ResponseBody[Rsp]) PrepareResponse(result *rrsp.ResponseResult) *errpkg.Error {
	rawResponse := result.GetRespBytes()
	if r.Body == nil {
		return nil
	}
	if err := json.Unmarshal(rawResponse, r.Body); err != nil {
		return errpkg.NewError(
			fmt.Errorf("unmarshal response: %w - '%v'", err, string(rawResponse)),
		)
	}
	return nil
}
