package rest

import (
	"net/http"

	"github.com/shonjord/tracker/internal/pkg/domain/errors"
)

type (
	httpHandler interface {
		Handle(req *Request, res *Response) error
	}
)

// toHttpError receives and error and maps any domain error to HTTP ones.
func toHttpError(err error) *Error {
	if httpErr, ok := err.(*Error); ok {
		return httpErr
	}

	if domainErr, ok := err.(*errors.ConflictError); ok {
		return &Error{
			Message:               domainErr.Message,
			HTTPStatusCounterpart: http.StatusConflict,
		}
	}
	if domainErr, ok := err.(*errors.NotFoundError); ok {
		return &Error{
			Message:               domainErr.Message,
			HTTPStatusCounterpart: http.StatusNotFound,
		}
	}
	if domainErr, ok := err.(*errors.InternalError); ok {
		return &Error{
			Message:               domainErr.Message,
			HTTPStatusCounterpart: http.StatusInternalServerError,
		}
	}

	return &Error{
		Message:               "internal server error",
		HTTPStatusCounterpart: http.StatusInternalServerError,
	}
}
