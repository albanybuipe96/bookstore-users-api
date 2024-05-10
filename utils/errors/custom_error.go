package errors

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
	"net/http"
	"strings"
)

// CustomError represents a custom error structure with a message, HTTP status code, and a reason.
// The reason represents the error type such internal_server_error, not_found_error, bad_request.
type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"status,omitempty"`
	Reason  string `json:"error"`
}

// NewCustomError creates and returns a pointer to a new CustomError
func NewCustomError() *CustomError {
	return &CustomError{}
}

// String returns the error message as a string.
func (err *CustomError) String() string {
	return err.Message
}

// Error implements the error interface by returning the error message as a string.
func (err *CustomError) Error() string {
	return err.Message
}

// ReportError returns the error code and a new CustomError instance with the same details.
func (err *CustomError) ReportError() (int, *CustomError) {
	return err.Code, &CustomError{
		Message: err.Message,
		Code:    err.Code,
		Reason:  err.Reason,
	}
}

// BadRequestError creates a CustomError with a bad request status code and a custom message.
func BadRequestError(message string) *CustomError {
	return &CustomError{
		Code:    http.StatusBadRequest,
		Message: message,
		Reason:  "bad_request",
	}
}

// InternalServerError creates a CustomError with an internal server error status code and a custom message.
func InternalServerError(msg string) *CustomError {
	return &CustomError{
		Code:    http.StatusInternalServerError,
		Message: msg,
		Reason:  "internal_server_error",
	}
}

// NotFoundError creates a CustomError with a not found status code and a custom message.
func NotFoundError(msg string) *CustomError {
	return &CustomError{
		Code:    http.StatusNotFound,
		Message: msg,
		Reason:  "not_found_error",
	}
}

// ReportDbError reports database errors, returning a CustomError based on the error type.
// It handles MySQL errors and specific cases like "no rows in result set" or duplicate entries.
func ReportDbError(err error) *CustomError {
	var sqlErr *pq.Error
	ok := errors.As(err, &sqlErr)
	if !ok {

		if strings.Contains(err.Error(), "no rows in result set") {
			return NotFoundError("no record found for given id")
		}

		return InternalServerError(err.Error())
	}

	fmt.Println("PG ERROR", sqlErr.Code, sqlErr.Message)

	// https://www.postgresql.org/docs/9.3/errcodes-appendix.html

	switch sqlErr.Code {
	case "42601":
		return InternalServerError("Syntax error in db")
	case "23505":
		return BadRequestError("Email already taken")

	}

	return InternalServerError("Something went wrong")
}
