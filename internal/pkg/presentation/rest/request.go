package rest

import (
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

type (
	Request struct {
		request *http.Request
	}
)

// NewRequest creates a new instance of Request struct.
func NewRequest(r *http.Request) *Request {
	return &Request{
		request: r,
	}
}

// GetParam returns the given param of this request.
func (r *Request) GetParam(p string) string {
	return chi.URLParam(r.request, p)
}

// HasQuery verifies if the URL has the given query param of this request.
func (r *Request) HasQuery(q string) bool {
	return "" != r.GetQuery(q)
}

// GetQuery returns the given query of this request.
func (r *Request) GetQuery(q string) string {
	return r.request.URL.Query().Get(q)
}

// Body returns the current body of this request.
func (r *Request) Body() io.ReadCloser {
	return r.request.Body
}

// HasHeader verifies if the given key exists in the headers.
func (r *Request) HasHeader(k string) bool {
	return "" != r.request.Header.Get(k)
}

// GetHeader gets the first value associated with the given key.
func (r *Request) GetHeader(k string) string {
	return r.request.Header.Get(k)
}
