package server

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/core-pb/tag/tag/v1/tagconnect"
	"github.com/uptrace/bun"
)

type Server struct {
	db *bun.DB
}

func New() *Server { return &Server{} }

func (s *Server) SetDB(ctx context.Context, db *bun.DB) *Server {
	var (
		_ctx, cancel = context.WithCancelCause(ctx)
		model        = []any{&Module{}, &Type{}, &Tag{}, &Relation{}}
	)

	db.RegisterModel(model...)
	for _, v := range model {
		if _, err := db.NewCreateTable().IfNotExists().Model(v).Exec(ctx); err != nil {
			cancel(err)
			break
		}
	}
	if err := _ctx.Err(); err != nil {
		slog.Error("sync model", slog.String("err", err.Error()))
		os.Exit(1)
	}

	s.db = db
	return s
}

func (s *Server) Register(handle func(pattern string, handler http.Handler)) *Server {
	handle(tagconnect.NewBaseHandler(base{Server: s}))
	handle(tagconnect.NewRelationshipHandler(relationship{Server: s}))
	handle(tagconnect.NewInternalHandler(itn{Server: s}))

	return s
}
