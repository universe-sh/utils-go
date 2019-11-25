package httputils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// Response HTTP
type Response struct {
	RequestID string      `json:"request_id,omitempty"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

var (
	errorCode = map[int]string{
		401: "Unauthorized",
		403: "Forbidden",
		404: "Not Found",
		409: "Conflict",
		500: "Internal Server Error",
		503: "Service Unavailable",
	}
	successCode = map[int]string{
		200: "OK",
		201: "Created",
		204: "No Content",
	}
)

// SendError response
func SendError(w http.ResponseWriter, code int) {
	var (
		guuid = uuid.New()
		err   error
	)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-ID", guuid.String())
	w.WriteHeader(code)

	err = json.NewEncoder(w).Encode(Response{
		RequestID: guuid.String(),
		Message:   errorCode[code],
	})
	if err != nil {
		log.Printf("RequestID %s: failed json %v", guuid.String(), err)
		return
	}
}

// SendData response
func SendData(w http.ResponseWriter, code int, data interface{}) {
	var (
		guuid = uuid.New()
		err   error
	)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-ID", guuid.String())
	w.WriteHeader(code)

	err = json.NewEncoder(w).Encode(Response{Data: data})
	if err != nil {
		log.Printf("RequestID %s: failed json %v", guuid.String(), err)
		return
	}
}

// SendCode response
func SendCode(w http.ResponseWriter, code int) {
	var (
		guuid = uuid.New()
		err   error
	)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-ID", guuid.String())
	w.WriteHeader(code)

	err = json.NewEncoder(w).Encode(Response{
		RequestID: guuid.String(),
		Message:   successCode[code],
	})
	if err != nil {
		log.Printf("RequestID %s: failed json %v", guuid.String(), err)
		return
	}
}
