package errors

type ErrorApp struct {
	Code    int
	Message string
}

func (err *ErrorApp) Error() string {
	return err.Message
}

func NotFoundError(message string) *ErrorApp {
	return &ErrorApp{
		Code:    404,
		Message: message,
	}
}

func BadRequestError(message string) *ErrorApp {
	return &ErrorApp{
		Code:    400,
		Message: message,
	}
}
