package httputils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

// Response HTTP
type Response struct {
	RequestID string        `json:"request_id,omitempty"`
	Message   string        `json:"message,omitempty"`
	Results   []interface{} `json:"results,omitempty"`
}

var (
	errorCode = map[int]string{
		401: "Unauthorized",
		403: "Forbidden",
		404: "Not Found",
		409: "Conflict",
		500: "Internal Server Error",
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

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(Response{
		RequestID: guuid.String(),
		Message:   errorCode[code],
	})
	if err != nil {
		log.Printf("Failed json: %v", err)
		return
	}

	return
}

// SendData response
func SendData(w http.ResponseWriter, code int, data []interface{}) {
	var err error

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(Response{Results: data})
	if err != nil {
		log.Printf("Failed json: %v", err)
		return
	}

	return
}

// SendCode response
func SendCode(w http.ResponseWriter, code int) {
	var (
		guuid = uuid.New()
		err   error
	)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(Response{
		RequestID: guuid.String(),
		Message:   successCode[code],
	})
	if err != nil {
		log.Printf("Failed json: %v", err)
		return
	}

	return
}
