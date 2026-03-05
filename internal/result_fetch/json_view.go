package result_fetch

import (
	"encoding/json"
)

func (r Result) String() string {
	b, err := json.Marshal(r)
	if err != nil {
		return ""
	}
	return string(b)
}
