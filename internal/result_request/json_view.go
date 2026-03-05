package result_request

import (
	"encoding/json"
)

type RequestResultJSON struct {
	IsEmpty bool   `json:"is_empty"`
	Method  string `json:"method"`
	URL     string `json:"url"`
	Curl    string `json:"curl"`
}

func (r RequestResult) ToJSON() RequestResultJSON {
	return RequestResultJSON{
		IsEmpty: r.IsEmpty(),
		Method:  r.GetMethod(),
		URL:     r.GetURL(),
		Curl:    r.GetCurl(),
	}
}

func (r RequestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.ToJSON())
}

func (r RequestResult) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(b)
}
