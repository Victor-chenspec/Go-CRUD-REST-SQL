package apperror

import (
	"net/http"
)

type AppError struct {
	Status  int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

func NotFound(message string) *AppError {
	return &AppError{
		Status: http.StatusNotFound,
		Message: message,
	}
}

func InternalError(message string) *AppError {
	return &AppError{
		Status: http.StatusInternalServerError,
		Message: message,
	}
}

func BadRequest(message string) *AppError {
	return &AppError{
		Status: http.StatusBadRequest,
		Message: message,
	}
}

func Unprocessable(message string) *AppError {
	return &AppError{
		Status: http.StatusUnprocessableEntity,
		Message: message,
	}
}