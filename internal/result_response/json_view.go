package result_response

import (
	"encoding/json"
)

type ResponseResultJSON struct {
	IsEmpty    bool   `json:"is_empty"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error,omitempty"`
	Body       string `json:"body,omitempty"`
}

func (r ResponseResult) ToJSON() ResponseResultJSON {
	var errStr string
	if e := r.GetHttpError(); e != nil {
		errStr = e.Error()
	}
	return ResponseResultJSON{
		IsEmpty:    r.IsEmpty(),
		StatusCode: r.GetStatusCode(),
		Error:      errStr,
		Body:       string(r.GetRespBytes()),
	}
}

func (r ResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.ToJSON())
}

func (r ResponseResult) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(b)
}
