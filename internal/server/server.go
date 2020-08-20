package server

import (
	"context"
	"errors"
	"net/http"
	"path"
	"sync"

	"github.com/gorilla/mux"
	"github.com/sedalu/interview/internal/store"
	"golang.org/x/sync/errgroup"
)

const (
	HeaderContentTypeKey   = "Content-Type"
	HeaderContentTypeValue = "application/json"
)

const (
	ParamKeyFib = "n"
)

const (
	EndpointFib     = "/fib"
	EndpointReverse = "/reverse"
	EndpointUnique  = "/unique"
)

type Server struct {
	store.Store

	r    *mux.Router
	once sync.Once
}

func (s *Server) handlers() {
	s.r.Methods(http.MethodGet).Path(path.Join(EndpointFib, "{"+ParamKeyFib+"}")).Handler(FibHandler{})
	s.r.Methods(http.MethodPost).Path(EndpointReverse).Handler(ReverseHandler{})
	s.r.Methods(http.MethodPost).Path(EndpointUnique).Handler(UniqueHandler{})
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.once.Do(s.handlers)

	w.Header().Set(HeaderContentTypeKey, HeaderContentTypeValue)

	s.r.ServeHTTP(w, r)
}

func (s *Server) Start(ctx context.Context) error {
	svr := &http.Server{
		Handler: s,
	}

	g := errgroup.Group{}

	g.Go(func() error {
		if err := svr.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				return err
			}
		}

		return nil
	})

	g.Go(func() error {
		<-ctx.Done()

		return svr.Shutdown(context.Background())
	})

	return g.Wait()
}
