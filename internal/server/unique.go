package server

import "net/http"

type UniqueHandler struct{}

// ServeHTTP responds to the request with a JSON array of unique values, preserving order, from the array provided in the request body.
func (h UniqueHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Get array of values from the request body.
	// TODO: Write unique values to the response body.
}
