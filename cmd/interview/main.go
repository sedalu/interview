package main

import (
	"context"
	"os"
	"time"

	"github.com/sedalu/interview/internal/server"
	"github.com/sedalu/interview/internal/store/pgsql"
)

func run(ctx context.Context) error {
	store, err := pgsql.Open(os.Getenv("DB_URL"))
	if err != nil {
		return err
	}

	svr := &server.Server{
		Store: store,
	}

	return svr.Start(ctx)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := run(ctx); err != nil {
		os.Exit(1)
	}
}
