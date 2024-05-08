package errors

import (
	"net/http"
)

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Reason  string `json:"error"`
}

func (err *CustomError) String() string {
	return err.Message
}

func (err *CustomError) Error() string {
	return err.Message
}

func (err *CustomError) ReportError() (int, *CustomError) {
	return err.Code, &CustomError{
		Message: err.Message,
		Code:    err.Code,
		Reason:  err.Reason,
	}
}

func BadRequestError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusBadRequest,
		Message: message,
		Reason:  "bad_request",
	}
}

func InternalServerError(msg string) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Reason:  "internal_server_error",
	}
}

func NotFoundError(msg string) *CustomError {
	return &CustomError{
		Code:    http.StatusNotFound,
		Message: msg,
		Reason:  "not_found_error",
	}
}
