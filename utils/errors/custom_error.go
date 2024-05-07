package errors

type CustomError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Error   string `error:"error"`
}

func (err *CustomError) ReportError() (int, CustomError) {
	return err.Code, CustomError{
		Message: err.Message,
		Code:    err.Code,
		Error:   err.Error,
	}
}
