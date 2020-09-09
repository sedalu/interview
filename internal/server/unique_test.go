package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUniqueHandler_ServerHTTP(t *testing.T) {
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
		{"unique1", request{`[1,1,2,4,3,5,4,4,1,2,6]`}, response{http.StatusOK, `[1,2,4,3,5,6]`}},
		{"unique2", request{`[9,7,5,3,1,2,4,6,8,0]`}, response{http.StatusOK, `[9,7,5,3,1,2,4,6,8,0]`}},
		{"unique3", request{`[5,3,3,3,1,4,2,2,0]`}, response{http.StatusOK, `[5,3,1,4,2,0]`}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodPost, EndpointReverse, strings.NewReader(tt.req.body))
			UniqueHandler{}.ServeHTTP(w, r)
			resp := w.Result()
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			if diff := cmp.Diff(tt.want.status, resp.StatusCode); diff != "" {
				t.Errorf("UniqueHandler.ServerHTTP() mismatch status (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want.body, strings.TrimSpace(string(body))); diff != "" {
				t.Errorf("UniqueHandler.ServerHTTP() mismatch body (-want +got):\n%s", diff)
			}
		})
	}
}
