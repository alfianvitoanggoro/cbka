package response

import (
	"encoding/json"
	"net/http"
)

type Meta struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Meta    *Meta  `json:"meta,omitempty"`
}

func SuccessResponse(message string, data any) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func SuccessResponseWithMeta(message string, data any, meta *Meta) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

func ErrorResponse(message string, err string) Response {
	return Response{
		Success: false,
		Message: message,
		Error:   err,
	}
}

// WriteSuccess writes a success response
func WriteSuccess(w http.ResponseWriter, code int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(SuccessResponse(message, data))
}

// WriteSuccessWithMeta writes a success response with pagination meta
func WriteSuccessWithMeta(w http.ResponseWriter, code int, message string, data any, meta *Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(SuccessResponseWithMeta(message, data, meta))
}

// WriteError writes an error response
func WriteError(w http.ResponseWriter, code int, message string, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(ErrorResponse(message, err))
}
