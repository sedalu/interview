package main

import (
	"context"
	"os"
	"time"

	"github.com/sedalu/interview/internal/server"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	svr := &server.Server{}

	if err := svr.Start(ctx); err != nil {
		os.Exit(1)
	}
}
