package iface

type ResponseBodyHandler interface {
	ResponseComponentHandler
	WithJsonObj(respObj any)
	WithJsonObjectWrapper(wrapper any, dataFieldName string, dataObj any)
}
