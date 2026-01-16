package smc

import (
	"fmt"
)

type StatusError struct {
	Code    int
	Message string
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("code %d, %s", e.Code, e.Message)
}

func NewStatusError(code int, message string) error {
	return &StatusError{
		Code:    code,
		Message: message,
	}
}
