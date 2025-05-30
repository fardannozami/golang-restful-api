package response

import (
	"encoding/json"
	"net/http"
)

// ApiResponse is the general structure for all API responses
type ApiResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func writeJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(payload)
}

func WriteError(w http.ResponseWriter, code int, message string) {
	writeJSON(w, code, ApiResponse{
		Code:    code,
		Message: message,
	})
}

func WriteValidationError(w http.ResponseWriter, message string) {
	WriteError(w, http.StatusBadRequest, message)
}

func WriteNotFoundError(w http.ResponseWriter, message string) {
	WriteError(w, http.StatusNotFound, message)
}

func WriteInternalError(w http.ResponseWriter, message string) {
	WriteError(w, http.StatusInternalServerError, message)
}

func WriteSuccess(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Message: message,
	})
}

func WriteCreated(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusCreated, ApiResponse{
		Code:    http.StatusCreated,
		Message: message,
	})
}

func WriteData(w http.ResponseWriter, data interface{}) {
	writeJSON(w, http.StatusOK, ApiResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	})
}
