package errors

import "net/http"

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `error:"error"`
}

func (err *CustomError) ReportError() (int, *CustomError) {
	return err.Code, &CustomError{
		Message: err.Message,
		Code:    err.Code,
		Error:   err.Error,
	}
}

func BadRequestError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusBadRequest,
		Message: message,
		Error:   "bad_request",
	}
}

func InternalServerError(msg string) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Error:   "internal_server_error",
	}
}
