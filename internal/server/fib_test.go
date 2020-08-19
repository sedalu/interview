package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"
)

func TestFibHandler_ServerHTTP(t *testing.T) {
	type request struct {
		n string
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
		{"fib1", request{"1"}, response{http.StatusOK, `[1]`}},
		{"fib2", request{"2"}, response{http.StatusOK, `[1,1]`}},
		{"fib3", request{"3"}, response{http.StatusOK, `[1,1,2]`}},
		{"fib4", request{"4"}, response{http.StatusOK, `[1,1,2,3]`}},
		{"fib5", request{"5"}, response{http.StatusOK, `[1,1,2,3,5]`}},
		{"fib6", request{"6"}, response{http.StatusOK, `[1,1,2,3,5,8]`}},
		{"fib7", request{"7"}, response{http.StatusOK, `[1,1,2,3,5,8,13]`}},
		{"fib8", request{"8"}, response{http.StatusOK, `[1,1,2,3,5,8,13,21]`}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			w, r := httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, path.Join(EndpointFib, tt.req.n), nil)
			FibHandler{}.ServeHTTP(w, mux.SetURLVars(r, map[string]string{ParamKeyFib: tt.req.n}))
			resp := w.Result()
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			if diff := cmp.Diff(tt.want.status, resp.StatusCode); diff != "" {
				t.Errorf("FibHandler.ServerHTTP() mismatch status (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want.body, strings.TrimSpace(string(body))); diff != "" {
				t.Errorf("FibHandler.ServerHTTP() mismatch body (-want +got):\n%s", diff)
			}
		})
	}
}
