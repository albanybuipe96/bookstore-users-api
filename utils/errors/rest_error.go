package errors

import (
	"net/http"
)

type RestError struct {
	Message string `json:"message"`
	Status    int    `json:"code"`
	Error   string `json:"error"`
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "Bad_Request",
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status: http.StatusNotFound,
		Error: "Not_Found",
	}
}