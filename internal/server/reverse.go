package server

import "net/http"

type ReverseHandler struct{}

// ServeHTTP responds to the request with a JSON string of reverse-ordered words from the request body.
func (h ReverseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: Get string from the request body.
	// TODO: Write reversed string to the response body.
}
