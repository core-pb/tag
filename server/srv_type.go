package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/core-pb/dt/time/v1"
	v1 "github.com/core-pb/tag/tag/v1"
	"github.com/uptrace/bun"
)

func (base) ListType(ctx context.Context, req *connect.Request[v1.ListTypeRequest]) (*connect.Response[v1.ListTypeResponse], error) {
	sq := db.NewSelect().Model(&Type{})
	sq = InOrEqPure(sq, `"type".id`, req.Msg.Id)
	sq = InOrEqPure(sq, `"type".key`, req.Msg.Key)
	sq = QueryFormStruct(sq, `"type".info`, req.Msg.Info)

	if req.Msg.Inherit != nil {
		sq = sq.Where(`"type".inherit = ?`, req.Msg.Inherit)
	}
	if req.Msg.Exclusive != nil {
		sq = sq.Where(`"type".exclusive = ?`, req.Msg.Exclusive)
	}

	if len(req.Msg.ModuleId) != 0 {
		const moduleVisibleType = `SELECT DISTINCT unnest(visible_type) AS type_id FROM "module" WHERE id IN (?)`
		sq = sq.Join(`INNER JOIN (`+moduleVisibleType+`) AS mvt ON mvt.type_id = "type".id`, bun.In(req.Msg.ModuleId))
	}

	sq = Pagination(sq, req.Msg.Pagination)
	sq = Sorts(sq, req.Msg.Sort)

	var (
		arr        []*v1.Type
		count, err = sq.ScanAndCount(ctx, &arr)
	)
	if err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.ListTypeResponse{Data: arr, Count: int64(count)}), nil
}

func (base) SetType(ctx context.Context, req *connect.Request[v1.SetTypeRequest]) (*connect.Response[v1.SetTypeResponse], error) {
	if !isInvalidKey(req.Msg.Key) {
		return nil, connect.NewError(connect.CodeInvalidArgument, errInvalidKey)
	}

	var val Type
	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if req.Msg.Id != 0 {
			_, err := tx.NewUpdate().Returning("*").Model(&val).Where("id = ?", req.Msg.Key).Set(`"key" = ?`, req.Msg.Key).Exec(ctx)
			return err
		}

		if err := tx.NewSelect().Model(&val).Where(`"key" = ?`, req.Msg.Key).Scan(ctx); err == nil || !errors.Is(err, sql.ErrNoRows) {
			return err
		}

		val.Type = &v1.Type{Key: req.Msg.Key}
		_, err := tx.NewInsert().Model(&val).Exec(ctx)
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetTypeResponse{Data: val.Type}), nil
}

func (base) SetTypeInfo(ctx context.Context, req *connect.Request[v1.SetTypeInfoRequest]) (*connect.Response[v1.SetTypeInfoResponse], error) {
	var val Type
	if _, err := db.NewUpdate().Model(&val).Where("id = ?", req.Msg.Id).Set("info = ?", req.Msg.Info).Returning("*").Exec(ctx); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.SetTypeInfoResponse{Data: val.Type}), nil
}

func (base) DeleteType(ctx context.Context, req *connect.Request[v1.DeleteTypeRequest]) (*connect.Response[v1.DeleteTypeResponse], error) {
	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if has, err := tx.NewSelect().Model(&Tag{}).Where("type_id IN (?)", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("tag has use type")
		}

		t := time.Now()
		_, err := tx.NewUpdate().Model(&Type{}).Where("id IN (?)", bun.In(req.Msg.Id)).
			Set(`deleted_at = ?, "key" = CONCAT(?, "key")`, t, fmt.Sprintf("::%d:", t.Seconds)).
			Exec(ctx)
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.DeleteTypeResponse{}), nil
}

func (base) UpdateTypeInherit(ctx context.Context, req *connect.Request[v1.UpdateTypeInheritRequest]) (*connect.Response[v1.UpdateTypeInheritResponse], error) {
	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if has, err := tx.NewSelect().Model(&Tag{}).Where("type_id IN (?)", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("tag has use type")
		}

		if has, err := tx.NewSelect().Model(&Tag{}).Join(`INNER JOIN "relation" ON tag.id = relation.tag_id`).
			Where("type_id IN (?)", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("type's tag has relation")
		}

		_, err := tx.NewUpdate().Model(&Type{}).Where("id IN (?)", bun.In(req.Msg.Id)).Set("inherit = ?", req.Msg.Inherit).Exec(ctx)
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.UpdateTypeInheritResponse{}), nil
}

func (base) UpdateTypeExclusive(ctx context.Context, req *connect.Request[v1.UpdateTypeExclusiveRequest]) (*connect.Response[v1.UpdateTypeExclusiveResponse], error) {
	if err := db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
		if has, err := tx.NewSelect().Model(&Tag{}).Where("type_id IN (?)", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("tag has use type")
		}

		if has, err := tx.NewSelect().Model(&Tag{}).Join(`INNER JOIN "relation" ON tag.id = relation.tag_id`).
			Where("type_id IN (?)", bun.In(req.Msg.Id)).Exists(ctx); err != nil {
			return err
		} else if has {
			return errors.New("type's tag has relation")
		}

		_, err := tx.NewUpdate().Model(&Type{}).Where("id IN (?)", bun.In(req.Msg.Id)).Set("exclusive = ?", req.Msg.Exclusive).Exec(ctx)
		return err
	}); err != nil {
		return nil, connect.NewError(connect.CodeUnavailable, err)
	}

	return connect.NewResponse(&v1.UpdateTypeExclusiveResponse{}), nil
}
