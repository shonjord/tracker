package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/shonjord/tracker/internal/pkg/domain/entity"
)

type (
	httpPoster interface {
		Post(string, io.Reader, map[string]string) (*http.Response, error)
	}

	AdminNotificationClient struct {
		url    string
		poster httpPoster
	}
)

// NewAdminNotificationClient returns a new instance of this HTTP client.
func NewAdminNotificationClient(u string, p httpPoster) *AdminNotificationClient {
	return &AdminNotificationClient{
		url:    u,
		poster: p,
	}
}

// Notify notifies to the client.
func (c *AdminNotificationClient) Notify(computer *entity.Computer) error {
	var (
		endpoint = fmt.Sprintf("%s/%s", c.url, "notify")
	)

	employee := computer.Employee
	abbreviation := "no abbreviation"
	if employee.HasAbbreviation() {
		abbreviation = employee.Abbreviation
	}

	request := struct {
		Level                string `json:"level"`
		EmployeeAbbreviation string `json:"employeeAbbreviation"`
		Message              string `json:"message"`
	}{
		Level:                "warning",
		EmployeeAbbreviation: abbreviation,
		Message: fmt.Sprintf(
			"employee %s has been assigned with more than 3 computers",
			employee.Name,
		),
	}

	res, err := c.post(request, endpoint)
	if err != nil {
		return err
	}
	defer closeResponseBody(res)

	if res.StatusCode == http.StatusOK {
		return nil

	}

	return err
}

// post is a reusable method to post data into the given endpoint.
func (c *AdminNotificationClient) post(r interface{}, endpoint string) (*http.Response, error) {
	body, err := toBodyFromStruct(r)
	if err != nil {
		return nil, err
	}

	resp, err := c.poster.Post(endpoint, body, c.defaultHeaders())
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// defaultHeaders returns the default headers of this client.
func (c *AdminNotificationClient) defaultHeaders() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}
}
