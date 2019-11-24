package httputils

import (
	"encoding/base64"
	"log"
	"net/http"
)

// AuthInfoHandler reads authentication info provided by the Endpoints proxy.
func AuthInfoHandler(w http.ResponseWriter, r *http.Request) []byte {
	encodedInfo := r.Header.Get("X-Endpoint-API-UserInfo")
	if encodedInfo == "" {
		SendError(w, http.StatusForbidden)
		return nil
	}

	b, err := base64.StdEncoding.DecodeString(encodedInfo)
	if err != nil {
		log.Println(err)
		SendError(w, http.StatusInternalServerError)
		return nil
	}

	return b
}
