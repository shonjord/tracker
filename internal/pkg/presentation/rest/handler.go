package rest

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Handler receives an action handler and executes it.
func Handler(h httpHandler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req := NewRequest(r)
		res := NewResponse(w)

		res.SetJSONContentType()

		if err := h.Handle(req, res); err != nil {
			httpError := toHttpError(err)

			res.WriteHeader(httpError.HTTPStatusCounterpart)

			if err = res.WriteHttpError(httpError); err != nil {
				log.WithError(err).Error("while writing HTTP error")
			}

			return
		}
	}
}
