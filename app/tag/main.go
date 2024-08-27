package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/core-pb/tag/tag/v1/tagconnect"
	"github.com/uptrace/bun"
	"go.x2ox.com/sorbifolia/bunpgd"
)

func main() {
	var (
		ctx, cancel = context.WithCancel(context.Background())
		closeCh     = make(chan os.Signal, 1)
		mux         = http.NewServeMux()
		srv         = &http.Server{
			BaseContext: func(listener net.Listener) context.Context { return ctx },
			Handler:     mux,
		}
	)

	mux.Handle(tagconnect.NewBaseHandler(base{}))
	mux.Handle(tagconnect.NewRelationshipHandler(relationship{}))

	initDB(ctx)

	go func() {
		err := srv.ListenAndServeTLS("./server.crt", "./server.key")

		switch {
		case errors.Is(err, http.ErrServerClosed):
			return
		case err != nil:
			closeCh <- fakeSignal{err: err}
		}
	}()

	signal.Notify(closeCh, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	sig := <-closeCh
	if e, is := sig.(fakeSignal); is {
		slog.Error("server exit", e.Attr())
		os.Exit(1)
		return
	}

	slog.Info("server exit signal: signal notify %s", slog.String("sig", sig.String()))

	cancel()

	if err := srv.Close(); err != nil {
		slog.Error("server close", slog.String("err", err.Error()))
	}
}

type fakeSignal struct{ err error }

func (s fakeSignal) Signal()         {}
func (s fakeSignal) String() string  { return fmt.Sprintf("err[fake-signal]: %s", s.err) }
func (s fakeSignal) Attr() slog.Attr { return slog.String("err", s.err.Error()) }

func initDB(ctx context.Context) {
	var err error
	if db, err = bunpgd.Open(os.Getenv("DB_DSN"),
		bunpgd.WithMaxOpenConns(256),
		bunpgd.WithMaxIdleConns(8),
		bunpgd.WithConnMaxIdleTime(time.Second*12),
		bun.WithDiscardUnknownColumns(),
	); err != nil {
		slog.Error("connect db", slog.String("err", err.Error()))
		os.Exit(1)
	}

	model := []any{&Tag{}, &Type{}, &Module{}, &Relation{}}
	db.RegisterModel(model...)

	for _, v := range model {
		if _, err = db.NewCreateTable().IfNotExists().Model(v).Exec(ctx); err != nil {
			slog.Error("create table", slog.String("err", err.Error()))
			os.Exit(1)
		}
	}
}
