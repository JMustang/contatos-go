package errorhandler

import (
	"net/http"
)

type HandlerErr struct {
	// Error message.
	Message string `json:"message" example:"error trying to process request"`

	// Error description.
	Err string `json:"error" example:"internal_server_error"`

	// Error code.
	Code int `json:"code" example:"500"`

	// Error causes.
	Causes []Causes `json:"causes"`
}

type Causes struct {
	/*
		Field is the name of the field that caused the error.
		@json
		@jsonTag field
	*/
	Field string `json:"field" example:"name"`

	/*
		Error message describing the cause.
		@json
		@jsonTag message
	*/
	Message string `json:"message" example:"name is required"`
}

func (r *HandlerErr) Error() string {
	return r.Message
}

func NewHandlerErr(message string, err string, code int, causes []Causes) *HandlerErr {
	return &HandlerErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *HandlerErr {
	return &HandlerErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *HandlerErr {
	return &HandlerErr{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *HandlerErr {
	return &HandlerErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *HandlerErr {
	return &HandlerErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *HandlerErr {
	return &HandlerErr{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *HandlerErr {
	return &HandlerErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
