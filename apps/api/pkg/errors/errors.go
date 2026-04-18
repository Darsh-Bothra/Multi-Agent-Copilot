package errors

type AppError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}



func NewBadRequest(msg string) AppError {
	return AppError{Message: msg, Code: 400}
}