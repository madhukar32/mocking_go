package httputil

import (
	"encoding/json"
	"net/http"
)

// marshalJSON encodes v as JSON, panicking if an error occurs.
func marshalJSON(v interface{}) []byte {
	bytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return bytes
}

// SendJSON encodes v as JSON and writes it to the response body. Panics
// if an encoding error occurs.
func SendJSON(w http.ResponseWriter, status int, v interface{}) {
	body := marshalJSON(v)
	w.Header().Set("content-type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	w.Write(body)
	w.Write([]byte("\n"))
}
