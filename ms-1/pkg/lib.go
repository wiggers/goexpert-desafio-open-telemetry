package lib

type AppError struct {
	ErrorMsg error
	Message  string
	Code     int
}

func (a *AppError) Error() string {
	return a.ErrorMsg.Error()
}

func NewAppError(errorMsg error, message string, code int) *AppError {
	return &AppError{
		ErrorMsg: errorMsg,
		Message:  message,
		Code:     code,
	}
}

type Response struct {
	Message interface{}
	Code    int
}
