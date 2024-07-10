package core

import "fmt"

func NewValidationErr(msg string) error {
	return &ErrValidation{
		Message: msg,
	}
}

type ErrValidation struct {
	Message string `json:"message"`
}

func (e *ErrValidation) Error() string {
	return e.Message
}

func ErrRequired(field string) string {
	return fmt.Sprintf("%s is required", field)
}

func ErrMinLength(field string, min string) string {
	return fmt.Sprintf("%s must have a minimum of %s", field, min)
}

func ErrMaxLength(field string, max string) string {
	return fmt.Sprintf("%s must have a maximum of %s", field, max)
}
