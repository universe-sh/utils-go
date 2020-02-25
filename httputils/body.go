package httputils

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/golang/gddo/httputil/header"
)

// DecodeJSONBody in http request
func DecodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			return errors.New("Content-Type header is not application/json")
		}
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1048576)

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&dst); err != nil {
		return err
	}

	dec.DisallowUnknownFields()
	if dec.More() {
		return errors.New("Request body must only contain a single JSON object")
	}

	dec.Buffered()

	return nil
}
