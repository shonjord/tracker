package rest

import (
	"encoding/json"
	"net/http"
)

type (
	Response struct {
		response http.ResponseWriter
	}
)

// NewResponse returns a new instance of Response.
func NewResponse(r http.ResponseWriter) *Response {
	return &Response{
		response: r,
	}
}

// WriteHeader sends an HTTP response header with the provided status code.
func (r *Response) WriteHeader(statusCode int) {
	r.response.WriteHeader(statusCode)
}

// WriteHttpError receives a struct and writes it to the connection as part of an HTTP reply.
func (r *Response) WriteHttpError(err *Error) error {
	resp, _ := json.Marshal(err)

	if err := r.WriteBytes(resp); err != nil {
		return err
	}

	return nil
}

// WriteStruct receives a struct and writes it to the connection as part of an HTTP reply.
func (r *Response) WriteStruct(s interface{}) error {
	resp, err := json.Marshal(s)
	if err != nil {
		return r.InternalServerError(err)
	}

	if err = r.WriteBytes(resp); err != nil {
		return err
	}

	return nil
}

// WriteBytes writes content(bytes) to the connection as part of an HTTP reply.
func (r *Response) WriteBytes(content []byte) error {
	_, err := r.response.Write(content)

	if err != nil {
		return err
	}

	return nil
}

// BadRequest indicates that the server cannot or will not process the request.
func (r *Response) BadRequest(err error) *Error {
	return &Error{
		HTTPStatusCounterpart: http.StatusBadRequest,
		Message:               err.Error(),
	}
}

// InternalServerError indicates that the server encountered an errors condition.
func (r *Response) InternalServerError(err error) *Error {
	return &Error{
		HTTPStatusCounterpart: http.StatusInternalServerError,
		Message:               err.Error(),
	}
}

// SetJSONContentType set application/json header to the response.
func (r *Response) SetJSONContentType() {
	r.response.Header().Set("Content-Type", "application/json")
}
