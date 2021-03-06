package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ErrorTypeUnprocessableEntity gin.ErrorType = 1 << 13
)

var (
	// ErrTetapTenangTetapSemangat custom error on unexpected error
	ErrTetapTenangTetapSemangat = CustomError{
		Message:  "Tetap Tenang Tetap Semangat",
		HTTPCode: http.StatusInternalServerError,
	}

	ErrUnauthorized = CustomError{
		Message:  "Unauthorized",
		HTTPCode: http.StatusUnauthorized,
	}

	ErrForbidden = CustomError{
		Message:  "Forbidden",
		HTTPCode: http.StatusForbidden,
	}

	ErrNotFound = CustomError{
		Message:  "Record not exist",
		HTTPCode: http.StatusNotFound,
	}

	ErrUnprocessableEntity = CustomError{
		Message:  "Unprocessable Entity",
		HTTPCode: http.StatusUnprocessableEntity,
	}
)

// CustomError holds data for customized error
type CustomError struct {
	Message  string `json:"message"`
	HTTPCode int    `json:"code"`
}

// Error is a function to convert error to string.
// It exists to satisfy error interface
func (c CustomError) Error() string {
	return c.Message
}
