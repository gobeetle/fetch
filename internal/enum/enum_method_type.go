package enum

type MethodType string

var EnumMethodType = struct {
	Get     MethodType
	Post    MethodType
	Put     MethodType
	Delete  MethodType
	Patch   MethodType
	Head    MethodType
	Options MethodType
	Connect MethodType
	Trace   MethodType
}{
	Get:     "GET",
	Post:    "POST",
	Put:     "PUT",
	Delete:  "DELETE",
	Patch:   "PATCH",
	Head:    "HEAD",
	Options: "OPTIONS",
	Connect: "CONNECT",
	Trace:   "TRACE",
}
