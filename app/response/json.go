package response

import (
	"encoding/json"
	"io"
)

// JSON returns the v type
// as json response.
func JSON(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}
