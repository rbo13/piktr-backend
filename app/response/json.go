package response

import (
	"net/http"
)

// JSON returns the v type
// as a json response
// using the `json.NewEncoder()`
func JSON(w http.ResponseWriter, v []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(v)

	if err != nil {
		return err
	}
	return nil
}
