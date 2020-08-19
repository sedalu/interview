package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RespondWith sets the response status code and, if not nil, writes v to w as JSON.
func RespondWith(w http.ResponseWriter, status int, v interface{}) {
	w.WriteHeader(status)

	if v == nil {
		return
	}

	json.NewEncoder(w).Encode(v)
}

// Unmarshal decodes the request body as JSON and unmarshals the value into v.
func Unmarshal(r *http.Request, v interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	dec.UseNumber()

	return dec.Decode(v)
}

// URLParam returns the value of request URL parameter coresponding to key. If the value isn't set, then an empty string returned.
func URLParam(r *http.Request, key string) string {
	vars := mux.Vars(r)

	if vars == nil {
		return ""
	}

	return vars[key]
}
