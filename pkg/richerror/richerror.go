package richerror

import "errors"

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindNotFound
	KindForbidden
	KindUnexpected
)

type Op string

type RichError struct {
	kind         Kind
	wrappedError error
	message      string
	operation    Op
	meta         map[string]interface{}
}

func New(operation Op) RichError {
	return RichError{
		operation: operation,
	}
}

func (e RichError) Error() string {
	return e.message
}
func (e RichError) Unwrap() error {
	return e.wrappedError
}
func (e RichError) Meta() map[string]interface{} {
	return e.meta
}
func (e RichError) Operation() Op {
	return e.operation
}
func (e RichError) Kind() Kind {
	if e.kind == 0 {

		var err RichError
		ok := errors.As(e.wrappedError, &err)
		if !ok {
			return 0
		}
		return err.Kind()
	}
	return e.kind
}
func (e RichError) Message() string {
	if e.message == "" {

		var err RichError
		ok := errors.As(e.wrappedError, &err)
		if !ok {
			return ""
		}
		return err.Message()
	}
	return e.message
}

func (e RichError) WithKind(kind Kind) RichError {
	e.kind = kind
	return e
}
func (e RichError) WithMessage(message string) RichError {
	e.message = message
	return e
}
func (e RichError) WithMeta(meta map[string]interface{}) RichError {
	e.meta = meta
	return e
}
func (e RichError) WithOperation(operation Op) RichError {
	e.operation = operation
	return e
}
func (e RichError) WithError(err error) RichError {
	e.wrappedError = err
	return e
}
