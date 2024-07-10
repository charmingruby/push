package core

import "fmt"

func NewInternalErr(location string) error {
	return &ErrInternal{
		Message: fmt.Sprintf("internal error on %s", location),
	}
}

type ErrInternal struct {
	Message string `json:"message"`
}

func (e *ErrInternal) Error() string {
	return e.Message
}

func NewNotFoundErr(entity string) error {
	return &ErrNotFound{
		Message: fmt.Sprintf("%s not found", entity),
	}
}

type ErrNotFound struct {
	Message string `json:"message"`
}

func (e *ErrNotFound) Error() string {
	return e.Message
}
