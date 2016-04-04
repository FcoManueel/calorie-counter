package errors

import "fmt"

type AppError struct {
	Code        string `json:"errorCode"`
	Description string `json:"description"`
}

var errorMap map[string]int

const (
	BAD_REQUEST  = "BAD_REQUEST"
	UNAUTHORIZED = "UNAUTHORIZED"
	FORBIDDEN    = "FORBIDDEN"

	DATABASE_ERROR        = "DATABASE_ERROR"
	INTERNAL_SERVER_ERROR = "GENERAL_ERROR"
)

func init() {
	errorMap = map[string]int{
		BAD_REQUEST:  400,
		UNAUTHORIZED: 401,
		FORBIDDEN:    403,

		DATABASE_ERROR:        500,
		INTERNAL_SERVER_ERROR: 500,
	}
}

func New(code string, descriptionFmt string, a ...interface{}) AppError {
	return AppError{
		Code:        code,
		Description: fmt.Sprintf(descriptionFmt, a...),
	}
}

func (err AppError) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Description)
}

func (err AppError) HTTPCode() int {
	if httpCode, ok := errorMap[err.Code]; ok {
		return httpCode
	}
	return 500
}
