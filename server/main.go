package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/core-pb/tag/tag/v1/tagconnect"
	"github.com/uptrace/bun"
	"go.x2ox.com/sorbifolia/bunpgd"
	"go.x2ox.com/sorbifolia/crpc"
)

func main() {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		closeCh     = make(chan os.Signal, 1)
		server, err = crpc.NewServer(
			crpc.WithHealthAndMetrics(":80", ""),
			crpc.WithCertFromCheck("CERT", "cert", "../build/output/cert"),
			crpc.WithCORS(nil),
		)
	)
	if err != nil {
		slog.Error("create server", slog.String("err", err.Error()))
		os.Exit(1)
	}

	initDB(ctx)

	server.Handle(tagconnect.NewBaseHandler(base{}))
	server.Handle(tagconnect.NewRelationshipHandler(relationship{}))
	server.Handle(tagconnect.NewInternalHandler(itn{}))

	if err = server.Run(); err != nil {
		slog.Error("server run", slog.String("err", err.Error()))
		os.Exit(1)
	}

	signal.Notify(closeCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	server.SetReady()
	sig := <-closeCh
	server.SetNoReady("server close")

	slog.Info("server exit signal: signal notify %s", slog.String("sig", sig.String()))

	cancel()

	if err = server.Close(); err != nil {
		slog.Error("server close", slog.String("err", err.Error()))
	}
}

func initDB(ctx context.Context) {
	var (
		_ctx, cancel = context.WithCancelCause(ctx)
		err          error
	)

	if db, err = bunpgd.Open(os.Getenv("DB_DSN"),
		bunpgd.WithMaxOpenConns(256),
		bunpgd.WithMaxIdleConns(8),
		bunpgd.WithConnMaxIdleTime(time.Second*12),
		bun.WithDiscardUnknownColumns(),
		bunpgd.WithSLog(),
		bunpgd.WithCreateTable(_ctx, cancel, &Module{}, &Type{}, &Tag{}, &Relation{}),
	); err != nil {
		slog.Error("connect db", slog.String("err", err.Error()))
		os.Exit(1)
	}

	if err = _ctx.Err(); err != nil {
		slog.Error("create table", slog.String("err", err.Error()))
		os.Exit(1)
	}
}
