package iface

type ResponseRawBodyHandler interface {
	ResponseComponentHandler
	WithRawBodyBytes(bodyBytes []byte)
}
