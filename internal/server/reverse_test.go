package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReverseHandler_ServerHTTP(t *testing.T) {
	type request struct {
		body string
	}

	type response struct {
		status int
		body   string
	}

	tests := []struct {
		name string
		req  request
		want response
	}{
		{"rev1", request{`"the quick brown fox jumps over the lazy dog"`}, response{http.StatusOK, `"dog lazy the over jumps fox brown quick the"`}},
		{"rev2", request{`"the cow jumped over the moon"`}, response{http.StatusOK, `"moon the over jumped cow the"`}},
		{"rev3", request{`"See Spot run!"`}, response{http.StatusOK, `"run! Spot See"`}},
		{"rev4", request{`"Run! Spot Run!"`}, response{http.StatusOK, `"Run! Spot Run!"`}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, EndpointReverse, strings.NewReader(tt.req.body))
			ReverseHandler{}.ServeHTTP(w, r)
			resp := w.Result()
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			if diff := cmp.Diff(tt.want.status, resp.StatusCode); diff != "" {
				t.Errorf("ReverseHandler.ServerHTTP() mismatch status (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want.body, strings.TrimSpace(string(body))); diff != "" {
				t.Errorf("ReverseHandler.ServerHTTP() mismatch body (-want +got):\n%s", diff)
			}
		})
	}
}
