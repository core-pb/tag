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
	if x.Module == nil {
		x.Module = &tag.Module{}
	}

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
	if x.Tag == nil {
		x.Tag = &tag.Tag{}
	}

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
	if x.Relation == nil {
		x.Relation = &tag.Relation{}
	}

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
	if x.Type == nil {
		x.Type = &tag.Type{}
	}

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
	if _, err := query.NewCreateIndex().IfNotExists().Model((*Type)(nil)).Unique().
		Index("idx_type_unique_key").Column("key").Exec(ctx); err != nil {
		return err
	}
	if _, err := query.NewCreateIndex().IfNotExists().Model((*Type)(nil)).
		Index("idx_type_query").Column("info", "inherit", "exclusive").Exec(ctx); err != nil {
		return err
	}
	return nil
}

func (*Tag) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	if _, err := query.NewCreateIndex().IfNotExists().Model((*Tag)(nil)).Unique().
		Index("idx_tag_unique_key").Column("key").Exec(ctx); err != nil {
		return err
	}
	if _, err := query.NewCreateIndex().IfNotExists().Model((*Tag)(nil)).
		Index("idx_tag_query").Column("type_id", "parent_id", "data", "info").Exec(ctx); err != nil {
		return err
	}

	if err := foreignKeyIfNotExists(ctx, query.DB(), "tag", "type_id", "fk_tag_type_id", "type", "id"); err != nil {
		return err
	}

	return nil
}

func (*Relation) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	if _, err := query.NewCreateIndex().IfNotExists().Model((*Relation)(nil)).
		Index("idx_relation_query").Column("source_id", "data").Exec(ctx); err != nil {
		return err
	}

	if err := foreignKeyIfNotExists(ctx, query.DB(), "relation", "tag_id", "fk_rel_tag_id", "tag", "id"); err != nil {
		return err
	}
	if err := foreignKeyIfNotExists(ctx, query.DB(), "relation", "source_id", "fk_rel_tag_id_src", "tag", "id"); err != nil {
		return err
	}
	if err := foreignKeyIfNotExists(ctx, query.DB(), "relation", "module_id", "fk_rel_module_id", "module", "id"); err != nil {
		return err
	}

	return nil
}

func (*Module) AfterCreateTable(ctx context.Context, query *bun.CreateTableQuery) error {
	if _, err := query.NewCreateIndex().IfNotExists().Model((*Module)(nil)).Unique().
		Index("idx_module_unique_key").Column("key").Exec(ctx); err != nil {
		return err
	}
	if _, err := query.NewCreateIndex().IfNotExists().Model((*Module)(nil)).
		Index("idx_module_query").Column("info", "visible_type").Exec(ctx); err != nil {
		return err
	}
	return nil
}

func foreignKeyIfNotExists(ctx context.Context, db *bun.DB,
	table, key string, constraintName string,
	refTable, refKey string) error {

	if has, err := db.NewSelect().Table("information_schema.table_constraints").Where(
		"constraint_type = 'FOREIGN KEY' AND table_name = ? AND constraint_name = ?",
		table, constraintName,
	).Exists(ctx); err != nil || has {
		return err
	}
	if _, err := db.NewRaw(`ALTER TABLE ? ADD CONSTRAINT ? FOREIGN KEY (?) REFERENCES ? (?)`,
		bun.Name(table), bun.Name(constraintName), bun.Name(key), bun.Name(refTable), bun.Name(refKey),
	).Exec(ctx); err != nil {
		return err
	}

	return nil
}
