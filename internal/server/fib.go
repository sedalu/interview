package server

import (
	"net/http"
)

type FibHandler struct{}

// ServeHTTP responds to the request with a JSON array of the first n numbers of the fibonacci sequence.
func (h FibHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Get n from url parameter.
	// TODO: Write slice of fibonacci numbers to the response body.
}
