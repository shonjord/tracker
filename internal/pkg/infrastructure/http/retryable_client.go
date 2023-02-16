package http

import (
	"io"
	"net/http"
	"time"

	"github.com/rafaeljesus/retry-go"
)

const (
	retryAttempts = 3
	wait          = time.Second * 2
)

type (
	Retryable struct {
		inner *http.Client
	}
)

// NewRetryable returns a new instance of a client with retry mechanisms.
func NewRetryable() *Retryable {
	return &Retryable{
		inner: http.DefaultClient,
	}
}

// Post issues a POST to the specified URL.
func (c *Retryable) Post(url string, body io.Reader, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	if headers != nil {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	return doWithRetry(func() (*http.Response, error) {
		return c.inner.Do(req)
	})
}

// doWithRetry performs an HTTP operation with a retry mechanism.
func doWithRetry(r func() (*http.Response, error)) (*http.Response, error) {
	return retry.DoHTTP(r, retryAttempts, wait)
}
