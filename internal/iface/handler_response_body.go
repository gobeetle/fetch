package iface

type ResponseBodyHandler[Rsp any] interface {
	ResponseComponentHandler
	WithJsonObj(respObj *Rsp)
}
