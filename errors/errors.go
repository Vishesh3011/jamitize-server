package errors

import (
	"example/internal/types"
	"fmt"
)

type AppError interface {
	ErrorStr() string
	Error() error
	Status() types.Status
	Message() any
	Layer() types.Layer
	Json() map[string]any
}

func ToAppError(err error, status types.Status, layer types.Layer) AppError {
	if err == nil {
		return nil
	}

	return NewAppError(err, status, nil, layer)
}

type appError struct {
	cause   error
	status  types.Status
	message any
	layer   types.Layer
}

func (a appError) Error() error {
	return fmt.Errorf("%s", a.ErrorStr())
}

func (a appError) Json() map[string]any {
	return map[string]any{
		"status":  a.status,
		"message": a.message,
		"error":   a.cause,
		"layer":   a.layer,
	}
}

func (a appError) ErrorStr() string {
	if a.cause != nil {
		return fmt.Sprintf("status=%d, message=%v, cause=%v layer=%v", a.status, a.message, a.cause, a.layer)
	}
	return fmt.Sprintf("status=%d, message=%v, layer=%v", a.status, a.message, a.layer)
}

func (a appError) Status() types.Status {
	return a.status
}

func (a appError) Message() any {
	return a.message
}

func (a appError) Layer() types.Layer {
	return a.layer
}

func NewAppError(err error, status types.Status, message any, layer types.Layer) AppError {
	return appError{
		cause:   err,
		status:  status,
		message: message,
		layer:   layer,
	}
}
