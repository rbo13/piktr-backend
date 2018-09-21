package response

import (
	"encoding/json"
	"io"
)

// JSON returns the v type
// as a json response
// using the `json.NewEncoder()`
func JSON(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
