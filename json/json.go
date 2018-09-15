package json

import (
	"encoding/json"
	"io"
)

// ToJSON returns the v type as json.
func ToJSON(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
