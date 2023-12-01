package main

import (
	"context"
	"log"
	"time"

	"github.com/bersennaidoo/otelobserve/application/rest/server"
	"github.com/bersennaidoo/otelobserve/physical/otrace"
	"github.com/bersennaidoo/otelobserve/physical/zlog"
	"go.uber.org/zap"
)

func main() {

	tp, err := otrace.InitTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	logger := zlog.NewZlog("otelobserve")
	logger.Info(
		"failed to fetch URL",
		zap.String("url", "https://github.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	srv := server.NewServer()
	srv.Run()
}
