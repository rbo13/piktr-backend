package response

import (
	"encoding/json"
	"net/http"
)

// MIME types
const (
	MIMEApplicationJSON            = "application/json"
	MIMEApplicationJSONCharsetUTF8 = MIMEApplicationJSON + "; " + charsetUTF8
)

// character set
const (
	charsetUTF8 = "charset=UTF-8"
)

// JSON returns the v type
// as a json response
// by adding a content-type
// as application/json
func JSON(statusCode int, w http.ResponseWriter, v interface{}) error {
	b, err := json.Marshal(v)

	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", MIMEApplicationJSONCharsetUTF8)
	w.WriteHeader(statusCode)
	_, err = w.Write(b)

	if err != nil {
		return err
	}

	return nil
}
