package httperr

import (
	"fmt"
	"net/http"

	"github.com/nelsonfalves/boot-help/util"
)

type HttpError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
}

func (mr *HttpError) String() string {
	if mr == nil {
		return "message:"
	}
	s := fmt.Sprintf("message: %s", mr.Message)
	if !util.EmptyString(mr.Details) {
		s += ", details: " + mr.Details
	}
	return s
}

// Add details to HttpError
func (mr *HttpError) WithDetails(details string) *HttpError {
	if mr != nil {
		mr.Details = details
	}
	return mr
}

// Create a new HttpError with StatusCode 400 and a custom message
func Bad(message string) *HttpError {
	return &HttpError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

// Create a new HttpError with StatusCode 404 and a custom message
func NotFound(message string) *HttpError {
	return &HttpError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

// Create a new HttpError with StatusCode 409 and a custom message
func Conflict(message string) *HttpError {
	return &HttpError{
		StatusCode: http.StatusConflict,
		Message:    message,
	}
}

// Create a new HttpError with StatusCode 412 and a custom message
func Condition(message string) *HttpError {
	return &HttpError{
		StatusCode: http.StatusPreconditionFailed,
		Message:    message,
	}
}

// Create a new HttpError with StatusCode 500 and a custom message
func Internal(message string) *HttpError {
	return &HttpError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

func Unauthorized(message string) *HttpError {
	return &HttpError{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	}
}
