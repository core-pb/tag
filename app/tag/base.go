package main

import (
	"context"

	"github.com/core-pb/dt/time/v1"
	"github.com/core-pb/tag/tag/v1"
	"github.com/redis/rueidis"
	"github.com/uptrace/bun"
)

var (
	db    *bun.DB
	cache rueidis.Client
)

type Type struct {
	bun.BaseModel `bun:"table:type"`
	*tag.Type
}

type Module struct {
	bun.BaseModel `bun:"table:module"`
	*tag.Module
}

type Tag struct {
	bun.BaseModel `bun:"table:tag"`
	*tag.Tag
}

type Relation struct {
	bun.BaseModel `bun:"table:relation"`
	*tag.Relation
}

func (x *Module) BeforeAppendModel(_ context.Context, query bun.Query) error {
	t := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		if x.CreatedAt == nil {
			x.CreatedAt = t
		}
		if x.UpdatedAt == nil {
			x.UpdatedAt = t
		}
	case *bun.UpdateQuery:
		x.UpdatedAt = t
	}
	return nil
}

func (x *Tag) BeforeAppendModel(_ context.Context, query bun.Query) error {
	t := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		if x.CreatedAt == nil {
			x.CreatedAt = t
		}
		if x.UpdatedAt == nil {
			x.UpdatedAt = t
		}
	case *bun.UpdateQuery:
		x.UpdatedAt = t
	}
	return nil
}

func (x *Relation) BeforeAppendModel(_ context.Context, query bun.Query) error {
	t := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		if x.CreatedAt == nil {
			x.CreatedAt = t
		}
		if x.UpdatedAt == nil {
			x.UpdatedAt = t
		}
	case *bun.UpdateQuery:
		x.UpdatedAt = t
	}
	return nil
}

func (x *Type) BeforeAppendModel(_ context.Context, query bun.Query) error {
	t := time.Now()
	switch query.(type) {
	case *bun.InsertQuery:
		if x.CreatedAt == nil {
			x.CreatedAt = t
		}
		if x.UpdatedAt == nil {
			x.UpdatedAt = t
		}
	case *bun.UpdateQuery:
		x.UpdatedAt = t
	}
	return nil
}

func (*Type) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().IfNotExists().
		Model((*Type)(nil)).Index("idx_type_query").
		Column("info", "inherit", "exclusive").
		Exec(ctx)
	return err
}

func (*Tag) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().IfNotExists().
		Model((*Tag)(nil)).Index("idx_tag_query").
		Column("type_id", "parent_id", "data", "info").
		Exec(ctx)
	return err
}

func (*Relation) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().IfNotExists().
		Model((*Relation)(nil)).Index("idx_relation_query").
		Column("source_id", "data").
		Exec(ctx)
	return err
}

func (*Module) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	_, err := query.DB().NewCreateIndex().IfNotExists().
		Model((*Module)(nil)).Index("idx_module_query").
		Column("info", "visible_type").
		Exec(ctx)
	return err
}
