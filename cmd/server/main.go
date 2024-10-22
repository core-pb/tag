package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/core-pb/tag/server"
	"github.com/uptrace/bun"
	"go.x2ox.com/sorbifolia/bunpgd"
	"go.x2ox.com/sorbifolia/crpc"
)

func main() {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		closeCh     = make(chan os.Signal, 1)
		srv, err    = crpc.NewServer(
			crpc.WithHealthAndMetrics(":80", ""),
			crpc.WithCertFromCheck("CERT", "cert", "../build/output/cert"),
			crpc.WithCORS(nil),
		)
	)
	if err != nil {
		slog.Error("create server", slog.String("err", err.Error()))
		os.Exit(1)
	}

	server.New().SetDB(ctx, initDB()).Register(srv.Handle)

	if err = srv.Run(); err != nil {
		slog.Error("server run", slog.String("err", err.Error()))
		os.Exit(1)
	}

	signal.Notify(closeCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	srv.SetReady()
	sig := <-closeCh
	srv.SetNoReady("server close")

	slog.Info("server exit signal: signal notify %s", slog.String("sig", sig.String()))

	cancel()

	if err = srv.Close(); err != nil {
		slog.Error("server close", slog.String("err", err.Error()))
	}
}

func initDB() *bun.DB {
	db, err := bunpgd.Open(os.Getenv("DB_DSN"),
		bunpgd.WithMaxOpenConns(256),
		bunpgd.WithMaxIdleConns(8),
		bunpgd.WithConnMaxIdleTime(time.Second*12),
		bun.WithDiscardUnknownColumns(),
		bunpgd.WithSLog(),
	)
	if err != nil {
		slog.Error("connect db", slog.String("err", err.Error()))
		os.Exit(1)
	}

	return db
}
