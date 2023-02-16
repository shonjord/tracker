package http

import (
	"bytes"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// toBodyFromStruct returns a new body that can be read by an HTTP client.
func toBodyFromStruct(s interface{}) (*bytes.Reader, error) {
	buffer := new(bytes.Buffer)
	if err := json.NewEncoder(buffer).Encode(s); err != nil {
		return nil, err
	}

	return bytes.NewReader(buffer.Bytes()), nil
}

// closeResponseBody closes body response and logs in case of error.
func closeResponseBody(r *http.Response) {
	if err := r.Body.Close(); err != nil {
		log.WithError(err).Warn("closing response body was not possible")
	}
}
