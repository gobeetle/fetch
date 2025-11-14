package enum

type StatusCodeFilterType string

var EnumStatusCodeFilterType = struct {
	Include2XX StatusCodeFilterType
	Include3XX StatusCodeFilterType
	Include4XX StatusCodeFilterType
	Include5XX StatusCodeFilterType
	AllInvalid StatusCodeFilterType
	AllValid   StatusCodeFilterType
}{

	Include2XX: "Include2XX",
	Include3XX: "Include3XX",
	Include4XX: "Include4XX",
	Include5XX: "Include5XX",
	AllInvalid: "AllInvalid",
	AllValid:   "AllValid",
}
