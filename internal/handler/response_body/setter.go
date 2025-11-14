package response_body

import (
	"reflect"
)

func (r *ResponseBody) WithJsonObj(resp_obj any) {
	r.Body = resp_obj
}

func (r *ResponseBody) WithJsonObjectWrapper(wrapper any, data_field_name string, data_obj any) {
	wrapperValue := reflect.ValueOf(wrapper)
	if wrapperValue.Kind() == reflect.Ptr {
		wrapperValue = wrapperValue.Elem()
	}
	wrapperField := wrapperValue.FieldByName(data_field_name)
	if wrapperField.IsValid() && wrapperField.CanSet() {
		if wrapperField.Type().AssignableTo(reflect.TypeOf(data_obj)) {
			wrapperField.Set(reflect.ValueOf(data_obj))
		} else {
			// r.Err = fmt.Errorf("type mismatch: cannot assign %T to %T", data_obj, wrapperField.Interface())
		}
	}
	r.Body = wrapper
	r.WrapperDataFieldName = data_field_name
}
